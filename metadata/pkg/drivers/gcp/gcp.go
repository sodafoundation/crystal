// Copyright 2023 The SODA Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gcp

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	backendpb "github.com/opensds/crsytal/backend/proto"
	"github.com/opensds/crsytal/metadata/pkg/db"
	"github.com/opensds/crsytal/metadata/pkg/model"
	"github.com/opensds/crsytal/metadata/pkg/utils"
	pb "github.com/opensds/crsytal/metadata/proto"
	log "github.com/sirupsen/logrus"
	"strings"
	"sync"
	"time"
)

type S3Cred struct {
	Ak string
	Sk string
}

var ObjList = ObjectList

func (myc *S3Cred) IsExpired() bool {
	return false
}

func (myc *S3Cred) Retrieve() (credentials.Value, error) {
	cred := credentials.Value{AccessKeyID: myc.Ak, SecretAccessKey: myc.Sk}
	return cred, nil
}

type GcpAdapter struct {
	Backend *backendpb.BackendDetail
	Session *session.Session
}

func (ad *GcpAdapter) DownloadObject() {
	log.Info("Implement me (gcp) driver")
}

func GetHeadObject(svc s3iface.S3API, bucketName *string, obj *model.MetaObject) {
	meta, err := svc.HeadObject(&s3.HeadObjectInput{Bucket: bucketName, Key: &obj.ObjectName})
	if err != nil {
		log.Errorf("cannot perform head object on object %v in bucket %v. failed with error: %v", obj.ObjectName, *bucketName, err)
		return
	}

	if meta.ServerSideEncryption != nil {
		obj.ServerSideEncryption = *meta.ServerSideEncryption
	}

	if meta.ContentType != nil {
		obj.ObjectType = *meta.ContentType
	}
	if meta.Expiration != nil {
		expiresTime, err := time.Parse(time.RFC1123, *meta.Expiration)
		if err != nil {
			log.Errorf("unable to parse given string to time type. error: %v. skipping ExpiresDate field", err)
		} else {
			obj.ExpiresDate = &expiresTime
		}
	}
	if meta.ReplicationStatus != nil {
		obj.ReplicationStatus = *meta.ReplicationStatus
	}
	if meta.WebsiteRedirectLocation != nil {
		obj.RedirectLocation = *meta.WebsiteRedirectLocation
	}

	if meta.VersionId != nil {
		obj.VersionId = *meta.VersionId
	}

	metadata := map[string]string{}
	for key, val := range meta.Metadata {
		metadata[key] = *val
	}
	obj.Metadata = metadata
}

func ObjectList(svc s3iface.S3API, bucketName string) ([]*model.MetaObject, int64, error) {
	output, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: &bucketName})
	if err != nil {
		log.Errorf("unable to list objects in bucket %v. failed with error: %v", bucketName, err)
		return nil, 0, err
	}

	numObjects := len(output.Contents)
	var totSize int64
	objectArray := make([]*model.MetaObject, numObjects)
	for objIdx, object := range output.Contents {
		obj := &model.MetaObject{}
		objectArray[objIdx] = obj
		obj.LastModifiedDate = object.LastModified
		obj.ObjectName = *object.Key
		obj.Size = *object.Size
		totSize += obj.Size

		if object.StorageClass != nil {
			obj.StorageClass = *object.StorageClass
		}

		acl, err := svc.GetObjectAcl(&s3.GetObjectAclInput{Bucket: &bucketName, Key: &obj.ObjectName})
		if err != nil {
			log.Errorf("unable to get object Acl. failed with error: %v", err)
		} else {
			access := []*model.Access{}
			for _, grant := range acl.Grants {
				access = append(access, utils.AclMapper(grant))
			}
			obj.ObjectAcl = access
		}

		GetHeadObject(svc, &bucketName, obj)
	}

	return objectArray, totSize, nil
}

func GetBucketMeta(buckIdx int, bucket *s3.Bucket, svc s3iface.S3API, backendRegion *string, bucketArray []*model.MetaBucket, wg *sync.WaitGroup) {
	defer wg.Done()

	loc, err := svc.GetBucketLocation(&s3.GetBucketLocationInput{Bucket: bucket.Name})
	if err != nil {
		log.Errorf("unable to get bucket location. failed with error: %v", err)
		return
	}

	if strings.ToLower(*loc.LocationConstraint) != *backendRegion {
		return
	}

	buck := &model.MetaBucket{}
	buck.Name = *bucket.Name
	buck.CreationDate = bucket.CreationDate
	buck.Region = *loc.LocationConstraint

	objectArray, totalSize, err := ObjList(svc, *bucket.Name)

	if err != nil {
		return
	}

	buck.Objects = objectArray
	buck.NumberOfObjects = len(objectArray)
	buck.TotalSize = totalSize

	bucketArray[buckIdx] = buck

	tags, err := svc.GetBucketTagging(&s3.GetBucketTaggingInput{Bucket: bucket.Name})
	if err == nil {
		tagset := map[string]string{}
		for _, tag := range tags.TagSet {
			tagset[*tag.Key] = *tag.Value
		}
		buck.BucketTags = tagset
	} else if !strings.Contains(err.Error(), "NoSuchTagSet") {
		log.Errorf("unable to get bucket tags. failed with error: %v", err)
	}

	acl, err := svc.GetBucketAcl(&s3.GetBucketAclInput{Bucket: bucket.Name})
	if err != nil {
		log.Errorf("unable to get bucket Acl. failed with error: %v", err)
	} else {
		access := []*model.Access{}
		for _, grant := range acl.Grants {
			access = append(access, utils.AclMapper(grant))
		}
		buck.BucketAcl = access
	}
}

func BucketList(svc s3iface.S3API, backendRegion *string) ([]*model.MetaBucket, error) {

	output, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		log.Errorf("unable to list buckets. failed with error: %v", err)
		return nil, err
	}

	bucketArray := make([]*model.MetaBucket, len(output.Buckets))
	wg := sync.WaitGroup{}
	for idx, bucket := range output.Buckets {
		wg.Add(1)
		go GetBucketMeta(idx, bucket, svc, backendRegion, bucketArray, &wg)
	}
	wg.Wait()

	bucketArrayFiltered := []*model.MetaBucket{}
	for _, buck := range bucketArray {
		if buck != nil {
			bucketArrayFiltered = append(bucketArrayFiltered, buck)
		}
	}

	return bucketArrayFiltered, err
}

func (ad *GcpAdapter) SyncMetadata(ctx context.Context, in *pb.SyncMetadataRequest) error {

	svc := s3.New(ad.Session)
	backendRegion := ad.Session.Config.Region
	buckArr, err := BucketList(svc, backendRegion)
	if err != nil {
		log.Errorf("metadata collection for backend id: %v failed with error: %v", ad.Backend.Id, err)
		return err
	}

	metaBackend := model.MetaBackend{}
	metaBackend.Id = ad.Backend.Id
	metaBackend.BackendName = ad.Backend.Name
	metaBackend.Type = ad.Backend.Type
	metaBackend.Region = ad.Backend.Region
	metaBackend.Buckets = buckArr
	metaBackend.NumberOfBuckets = int32(len(buckArr))
	newContext := context.TODO()
	err = db.DbAdapter.CreateMetadata(newContext, metaBackend)

	return err
}
