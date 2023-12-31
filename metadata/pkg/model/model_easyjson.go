// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	time "time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC80ae7adDecodeGithubComOpensdsMultiCloudMetadataPkgModel(in *jlexer.Lexer, out *MetaObject) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.ObjectName = string(in.String())
		case "lastModifiedDate":
			if in.IsNull() {
				in.Skip()
				out.LastModifiedDate = nil
			} else {
				if out.LastModifiedDate == nil {
					out.LastModifiedDate = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastModifiedDate).UnmarshalJSON(data))
				}
			}
		case "size":
			out.Size = int64(in.Int64())
		case "serverSideEncryption":
			out.ServerSideEncryption = string(in.String())
		case "versionId":
			out.VersionId = string(in.String())
		case "storageClass":
			out.StorageClass = string(in.String())
		case "redirectLocation":
			out.RedirectLocation = string(in.String())
		case "replicationStatus":
			out.ReplicationStatus = string(in.String())
		case "expiresDate":
			if in.IsNull() {
				in.Skip()
				out.ExpiresDate = nil
			} else {
				if out.ExpiresDate == nil {
					out.ExpiresDate = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.ExpiresDate).UnmarshalJSON(data))
				}
			}
		case "objectAcl":
			if in.IsNull() {
				in.Skip()
				out.ObjectAcl = nil
			} else {
				in.Delim('[')
				if out.ObjectAcl == nil {
					if !in.IsDelim(']') {
						out.ObjectAcl = make([]*Access, 0, 8)
					} else {
						out.ObjectAcl = []*Access{}
					}
				} else {
					out.ObjectAcl = (out.ObjectAcl)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Access
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Access)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.ObjectAcl = append(out.ObjectAcl, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "objectTags":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.ObjectTags = make(map[string]string)
				} else {
					out.ObjectTags = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 string
					v2 = string(in.String())
					(out.ObjectTags)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		case "metadata":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Metadata = make(map[string]string)
				} else {
					out.Metadata = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 string
					v3 = string(in.String())
					(out.Metadata)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
		case "objectType":
			out.ObjectType = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeGithubComOpensdsMultiCloudMetadataPkgModel(out *jwriter.Writer, in MetaObject) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.ObjectName))
	}
	{
		const prefix string = ",\"lastModifiedDate\":"
		out.RawString(prefix)
		if in.LastModifiedDate == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.LastModifiedDate).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"size\":"
		out.RawString(prefix)
		out.Int64(int64(in.Size))
	}
	{
		const prefix string = ",\"serverSideEncryption\":"
		out.RawString(prefix)
		out.String(string(in.ServerSideEncryption))
	}
	if in.VersionId != "" {
		const prefix string = ",\"versionId\":"
		out.RawString(prefix)
		out.String(string(in.VersionId))
	}
	if in.StorageClass != "" {
		const prefix string = ",\"storageClass\":"
		out.RawString(prefix)
		out.String(string(in.StorageClass))
	}
	if in.RedirectLocation != "" {
		const prefix string = ",\"redirectLocation\":"
		out.RawString(prefix)
		out.String(string(in.RedirectLocation))
	}
	if in.ReplicationStatus != "" {
		const prefix string = ",\"replicationStatus\":"
		out.RawString(prefix)
		out.String(string(in.ReplicationStatus))
	}
	if in.ExpiresDate != nil {
		const prefix string = ",\"expiresDate\":"
		out.RawString(prefix)
		out.Raw((*in.ExpiresDate).MarshalJSON())
	}
	if len(in.ObjectAcl) != 0 {
		const prefix string = ",\"objectAcl\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v4, v5 := range in.ObjectAcl {
				if v4 > 0 {
					out.RawByte(',')
				}
				if v5 == nil {
					out.RawString("null")
				} else {
					(*v5).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.ObjectTags) != 0 {
		const prefix string = ",\"objectTags\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v6First := true
			for v6Name, v6Value := range in.ObjectTags {
				if v6First {
					v6First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v6Name))
				out.RawByte(':')
				out.String(string(v6Value))
			}
			out.RawByte('}')
		}
	}
	if len(in.Metadata) != 0 {
		const prefix string = ",\"metadata\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v7First := true
			for v7Name, v7Value := range in.Metadata {
				if v7First {
					v7First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v7Name))
				out.RawByte(':')
				out.String(string(v7Value))
			}
			out.RawByte('}')
		}
	}
	if in.ObjectType != "" {
		const prefix string = ",\"objectType\":"
		out.RawString(prefix)
		out.String(string(in.ObjectType))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MetaObject) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComOpensdsMultiCloudMetadataPkgModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MetaObject) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComOpensdsMultiCloudMetadataPkgModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MetaObject) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComOpensdsMultiCloudMetadataPkgModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MetaObject) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComOpensdsMultiCloudMetadataPkgModel(l, v)
}
func easyjsonC80ae7adDecodeGithubComOpensdsMultiCloudMetadataPkgModel1(in *jlexer.Lexer, out *MetaBucket) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "creationDate":
			if in.IsNull() {
				in.Skip()
				out.CreationDate = nil
			} else {
				if out.CreationDate == nil {
					out.CreationDate = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.CreationDate).UnmarshalJSON(data))
				}
			}
		case "name":
			out.Name = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "region":
			out.Region = string(in.String())
		case "bucketAcl":
			if in.IsNull() {
				in.Skip()
				out.BucketAcl = nil
			} else {
				in.Delim('[')
				if out.BucketAcl == nil {
					if !in.IsDelim(']') {
						out.BucketAcl = make([]*Access, 0, 8)
					} else {
						out.BucketAcl = []*Access{}
					}
				} else {
					out.BucketAcl = (out.BucketAcl)[:0]
				}
				for !in.IsDelim(']') {
					var v8 *Access
					if in.IsNull() {
						in.Skip()
						v8 = nil
					} else {
						if v8 == nil {
							v8 = new(Access)
						}
						(*v8).UnmarshalEasyJSON(in)
					}
					out.BucketAcl = append(out.BucketAcl, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "numberOfObjects":
			out.NumberOfObjects = int(in.Int())
		case "numberOfFilteredObjects":
			out.NumberOfFilteredObjects = int(in.Int())
		case "objects":
			if in.IsNull() {
				in.Skip()
				out.Objects = nil
			} else {
				in.Delim('[')
				if out.Objects == nil {
					if !in.IsDelim(']') {
						out.Objects = make([]*MetaObject, 0, 8)
					} else {
						out.Objects = []*MetaObject{}
					}
				} else {
					out.Objects = (out.Objects)[:0]
				}
				for !in.IsDelim(']') {
					var v9 *MetaObject
					if in.IsNull() {
						in.Skip()
						v9 = nil
					} else {
						if v9 == nil {
							v9 = new(MetaObject)
						}
						(*v9).UnmarshalEasyJSON(in)
					}
					out.Objects = append(out.Objects, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "totalSize":
			out.TotalSize = int64(in.Int64())
		case "filteredBucketSize":
			out.FilteredBucketSize = int64(in.Int64())
		case "bucketTags":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.BucketTags = make(map[string]string)
				} else {
					out.BucketTags = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v10 string
					v10 = string(in.String())
					(out.BucketTags)[key] = v10
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeGithubComOpensdsMultiCloudMetadataPkgModel1(out *jwriter.Writer, in MetaBucket) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CreationDate != nil {
		const prefix string = ",\"creationDate\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((*in.CreationDate).MarshalJSON())
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	if in.Region != "" {
		const prefix string = ",\"region\":"
		out.RawString(prefix)
		out.String(string(in.Region))
	}
	if len(in.BucketAcl) != 0 {
		const prefix string = ",\"bucketAcl\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v11, v12 := range in.BucketAcl {
				if v11 > 0 {
					out.RawByte(',')
				}
				if v12 == nil {
					out.RawString("null")
				} else {
					(*v12).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"numberOfObjects\":"
		out.RawString(prefix)
		out.Int(int(in.NumberOfObjects))
	}
	if in.NumberOfFilteredObjects != 0 {
		const prefix string = ",\"numberOfFilteredObjects\":"
		out.RawString(prefix)
		out.Int(int(in.NumberOfFilteredObjects))
	}
	{
		const prefix string = ",\"objects\":"
		out.RawString(prefix)
		if in.Objects == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v13, v14 := range in.Objects {
				if v13 > 0 {
					out.RawByte(',')
				}
				if v14 == nil {
					out.RawString("null")
				} else {
					(*v14).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"totalSize\":"
		out.RawString(prefix)
		out.Int64(int64(in.TotalSize))
	}
	if in.FilteredBucketSize != 0 {
		const prefix string = ",\"filteredBucketSize\":"
		out.RawString(prefix)
		out.Int64(int64(in.FilteredBucketSize))
	}
	if len(in.BucketTags) != 0 {
		const prefix string = ",\"bucketTags\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v15First := true
			for v15Name, v15Value := range in.BucketTags {
				if v15First {
					v15First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v15Name))
				out.RawByte(':')
				out.String(string(v15Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MetaBucket) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComOpensdsMultiCloudMetadataPkgModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MetaBucket) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComOpensdsMultiCloudMetadataPkgModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MetaBucket) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComOpensdsMultiCloudMetadataPkgModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MetaBucket) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComOpensdsMultiCloudMetadataPkgModel1(l, v)
}
func easyjsonC80ae7adDecodeGithubComOpensdsMultiCloudMetadataPkgModel2(in *jlexer.Lexer, out *MetaBackend) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = string(in.String())
		case "backendName":
			out.BackendName = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "region":
			out.Region = string(in.String())
		case "buckets":
			if in.IsNull() {
				in.Skip()
				out.Buckets = nil
			} else {
				in.Delim('[')
				if out.Buckets == nil {
					if !in.IsDelim(']') {
						out.Buckets = make([]*MetaBucket, 0, 8)
					} else {
						out.Buckets = []*MetaBucket{}
					}
				} else {
					out.Buckets = (out.Buckets)[:0]
				}
				for !in.IsDelim(']') {
					var v16 *MetaBucket
					if in.IsNull() {
						in.Skip()
						v16 = nil
					} else {
						if v16 == nil {
							v16 = new(MetaBucket)
						}
						(*v16).UnmarshalEasyJSON(in)
					}
					out.Buckets = append(out.Buckets, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "numberOfBuckets":
			out.NumberOfBuckets = int32(in.Int32())
		case "numberOFilteredBuckets":
			out.NumberOfFilteredBuckets = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeGithubComOpensdsMultiCloudMetadataPkgModel2(out *jwriter.Writer, in MetaBackend) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"backendName\":"
		out.RawString(prefix)
		out.String(string(in.BackendName))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"region\":"
		out.RawString(prefix)
		out.String(string(in.Region))
	}
	{
		const prefix string = ",\"buckets\":"
		out.RawString(prefix)
		if in.Buckets == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v17, v18 := range in.Buckets {
				if v17 > 0 {
					out.RawByte(',')
				}
				if v18 == nil {
					out.RawString("null")
				} else {
					(*v18).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"numberOfBuckets\":"
		out.RawString(prefix)
		out.Int32(int32(in.NumberOfBuckets))
	}
	{
		const prefix string = ",\"numberOFilteredBuckets\":"
		out.RawString(prefix)
		out.Int32(int32(in.NumberOfFilteredBuckets))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MetaBackend) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComOpensdsMultiCloudMetadataPkgModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MetaBackend) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComOpensdsMultiCloudMetadataPkgModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MetaBackend) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComOpensdsMultiCloudMetadataPkgModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MetaBackend) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComOpensdsMultiCloudMetadataPkgModel2(l, v)
}
func easyjsonC80ae7adDecodeGithubComOpensdsMultiCloudMetadataPkgModel3(in *jlexer.Lexer, out *ListMetaDataResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(ListMetaDataResponse, 0, 8)
			} else {
				*out = ListMetaDataResponse{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v19 *MetaBackend
			if in.IsNull() {
				in.Skip()
				v19 = nil
			} else {
				if v19 == nil {
					v19 = new(MetaBackend)
				}
				(*v19).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v19)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeGithubComOpensdsMultiCloudMetadataPkgModel3(out *jwriter.Writer, in ListMetaDataResponse) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v20, v21 := range in {
			if v20 > 0 {
				out.RawByte(',')
			}
			if v21 == nil {
				out.RawString("null")
			} else {
				(*v21).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v ListMetaDataResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComOpensdsMultiCloudMetadataPkgModel3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ListMetaDataResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComOpensdsMultiCloudMetadataPkgModel3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ListMetaDataResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComOpensdsMultiCloudMetadataPkgModel3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ListMetaDataResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComOpensdsMultiCloudMetadataPkgModel3(l, v)
}
func easyjsonC80ae7adDecodeGithubComOpensdsMultiCloudMetadataPkgModel4(in *jlexer.Lexer, out *Access) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "displayName":
			out.DisplayName = string(in.String())
		case "emailAddress":
			out.EmailAddress = string(in.String())
		case "id":
			out.ID = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "uri":
			out.URI = string(in.String())
		case "permission":
			out.Permission = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeGithubComOpensdsMultiCloudMetadataPkgModel4(out *jwriter.Writer, in Access) {
	out.RawByte('{')
	first := true
	_ = first
	if in.DisplayName != "" {
		const prefix string = ",\"displayName\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.DisplayName))
	}
	if in.EmailAddress != "" {
		const prefix string = ",\"emailAddress\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.EmailAddress))
	}
	if in.ID != "" {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	if in.URI != "" {
		const prefix string = ",\"uri\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.URI))
	}
	if in.Permission != "" {
		const prefix string = ",\"permission\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Permission))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Access) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComOpensdsMultiCloudMetadataPkgModel4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Access) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComOpensdsMultiCloudMetadataPkgModel4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Access) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComOpensdsMultiCloudMetadataPkgModel4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Access) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComOpensdsMultiCloudMetadataPkgModel4(l, v)
}
