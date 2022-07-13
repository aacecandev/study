# Lesson 1 - Preparing your managed infrastructure

## Configure /etc/hosts in control node

192.168.1.26 control.example.com control
192.168.1.160 ansible29.example.com ansible29
192.168.1.161 ansible212.example.com ansible212
192.168.1.162 ubuntu.example.com ubuntu
192.168.1.163 windows.example.com windows
192.168.1.164 tower.example.com tower

## Setting up managed nodes

- Controller: Rocky Linux 8.6 with ansible 2.9 and openssh-server, sshpass installed
- Node: Ubuntu 20.04 LTS; SSH (openssh-server) is available and /etc/hosts name resolving set up

inventory

ubuntu  ansible_host=192.168.1.162

[rocky]
ansible29 ansible_host=192.168.1.160
ansible212 ansible_host=192.168.1.161

[FROM ROCKY 2.9] Set up Ubuntu with a user named ansible (assuming that it has a sudo user alejandro)

ansible -i inventory ubuntu -m user -a "name=ansible create_home=yes" -u alejandro -b -k -K

ubuntu | CHANGED => {
    "ansible_facts": {
        "discovered_interpreter_python": "/usr/bin/python3"
    },
    "changed": true,
    "comment": "",
    "create_home": true,
    "group": 1001,
    "home": "/home/ansible",
    "name": "ansible",
    "shell": "/bin/sh",
    "state": "present",
    "system": false,
    "uid": 1001
}

[FROM ROCKY 2.9] Set up Rocky with a user named ansible (assuming the Rocky machine has direct root access)

ansible -i inventory ansible29 -m user -a "name=ansible" -u root -k
  create_home=yes is default behavior in Rocky

[FROM ROCKY 2.9] Set up password for the user ansible

ansible -i inventory ubuntu -m shell -a "echo 'ansible:password' | chpasswd" -u alejandro -b -k -K
ubuntu | CHANGED | rc=0 >>

ansible -i inventory rocky -m shell -a "echo 'ansible:password' | passwd --stdin ansible" -u root -k

[FROM ROCKY 2.9] Using the local user ansible, create an SSH key pair and copy it to the managed nodes

ssh-keygen; ssh-copy-id ubuntu; ssh-copy-id ansible1

[FROM ROCKY 2.9] Test

ansible -i inventory all -m command -a "id" -u ansible

## Setting up privilege escalation

echo 'ansible ALL=(ALL) NOPASSWD:ALL' > /tmp/sudoers
sudo cp /tmp/sudoers /etc/sudoers.d/ansible
ansible rocky -i inventory -u root -k -m copy -a "src=/tmp/sudoers dest=/etc/sudoers.d/ansible"

ansible ubuntu -i inventory -u alejandro -k -b -K -m file -a "path=/etc/sudoers.d state=directory"
ansible ubuntu -i inventory -u alejandro -k -b -K -m copy -a "src=/tmp/sudoers dest=/etc/sudoers.d/ansible"

ansible -i inventory all -m command -a "ls -l /root" -b

## Setting up Windows as a managed host

1. Install Windows 2019 Server Standard (180 day eval version will do), ensuring there is a Windows user with admin privileges.
2. Log in as user with Administrator privileges.
3. Open powershell `winrm quickconfig`
4. Disable Internet Explorer Enhanced Security Configuration in Server Manager
5. Set up WinRM using the script provided on https://docs.ansible.com/ansible/latest/user_guide/windows_setup.html

If there's not a HTTPS Listener because there's no cert, use this
https://www.visualstudiogeeks.com/devops/how-to-configure-winrm-for-https-manually

- get hostname with `hostname`
- `New-SelfSignedCertificate -DnsName "<HOSTNAME>" -CertStoreLocation Cert:\LocalMachine\My`
- `New-SelfSignedCertificate -DnsName "WIN-MHBLQTSSBPK" -CertStoreLocation Cert:\LocalMachine\My`
- $thumbprint = (Get-ChildItem -Path Cert:\LocalMachine\My | Where-Object {$_.Subject -Match ">HOSTNAME>"}).Thumbprint

$selector_set = @{
  Address = "*"
  Transport = "HTTPS"
}
$value_set = @{
  CertificateThumbprint = $thumbprint
}
New-WSManInstance -ResourceURI "winrm/config/Listener" -SelectorSet $selector_set -ValueSet $value_set

6. Verify using `winrm enumerate winrm/config/Listener` and verify that http and https listeners are available

7. Control Panel > User Accounts > Manage Another Account > Add a user account ; Change the account type to Administrator

### Configuring authentication

Windows supports different authentication protocols.
If authentication has been configured to a strict protocol on Windows, you may have to set up Ansible accordingly:

- Basic: only supports local accounts - do not use
- Certificate: only supports local accounts - do not use
- Kerberos: only works for AD accounts
- NTLM: works for local accounts and AD accounts
- CredSSP: offers best support for local and AD accounts

To check the available authentication protocols, use `winrm get winrm/config/winrs` and `winrm get winrm/config/Service`. Lots of info and examples in the [ansible docs](https://docs.ansible.com/ansible/latest/user_guide/windows_setup.html)

Use `ansible_winrm_transport: credssp` (or anything else like `ansible_winrm_transport: ntlm`) if you have authentication issues

### Preparing the Control host

- In the Ansible controller, create a windows project folder with an ansible.cfg and an inventory (example in https://github.com/sandervanvugt/ansiblecvc)

```bash
# inventory

[win]
windows.example.com

[win:vars]
ansible_user=ansible
ansible_password=password
ansible_connection=winrm
ansible_winrm_server_cert_validation=ignore
```

- Set up /etc/hosts for hostname resolution to the Windows VM - notice that Windows firewall disallows ping incoming
- `sudo pip3 install pywinrm`
- `ansible win -i inventory -m ansible.windows.win_ping`
