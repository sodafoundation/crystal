# crsytal Installation Guide

## Ansible Installer
* Supported OS: **Ubuntu 18.04, 20.04**.
* Prerequisite: **Python 3.6+** and **Go 1.17+**  should be installed.

**Note**: Ensure root user access while performing the installation.

### Install Steps
```bash
apt-get update && apt-get install -y git
git clone https://github.com/sodafoundation/crsytal.git
cd crsytal/installer
chmod +x install_dependencies.sh && . install_dependencies.sh
```

### Set Host IP
```bash
cd ansible
export HOST_IP=<your_host_ip>
```

### Run the Ansible Playbook to install crsytal
```bash
ansible-playbook site.yml -i local.hosts
```
Ensure that the playbook finishes the installation without any failed steps. 

In case of any failure, run the following command to clear the installation:
```bash
ansible-playbook clean.yml -i local.hosts
```
Now, retry again:
```bash
ansible-playbook site.yml -i local.hosts
```

### Verify installation of crsytal
Run the below command to list the containers running:
```bash
docker ps
```
![Screenshot from 2023-04-13 10-34-50](https://user-images.githubusercontent.com/45416272/231659142-f8f1153b-c77a-4874-82c5-7d97510a13c0.png)

Ensure that all the above mentioned containers are up and running successfully. This confirms that crsytal has been installed successfully.

### Uninstall and Purge
```bash
ansible-playbook clean.yml -i local.hosts
```
### Using crsytal with dashboard
* Install crsytal with dashboard by following the installation guide [here](https://github.com/sodafoundation/installer/blob/master/README.md).
* For experiencing the features of crsytal, follow the user guide [here](https://docs.sodafoundation.io/guides/user-guides/crsytal/).

If you are interested in developing, follow the developer guide [here](https://docs.sodafoundation.io/guides/developer-guides/crsytal/).
