# Lesson 2 - Using Ad-hoc commands

## Ad-hoc commands vs playbooks

In an ad-hoc command, Ansible modules are used to perform tasks on managed hosts
Ad-hoc commands are perfect for running one task on a managed host:

- Initial host setup
- Checking configuration on hosts

Ad-hoc commands are also perfect to learn how easy Ansible really is to use.

Playbooks are the way to implement more complex tasks where dependency relations have to be managed.

### Understanding Ansible modules

An Ansible module is a Python script that will be executed in the managed host.

More than 3000 modules are available on Ansible.

Modules can be offered as such (ansible < 2.10) or as part of a collection (ansible >= 2.10).

### Running Ad-hoc commands

Use the ansible command to run ad-hoc commands

While running the ansible command, an inventory must be present, and you must specify the hosts to run the command on..

You'll also use the -m option to specify the module to use and -a option to specify the module arguments.

### Examples

This will fail as it will reboot and then request status

    ansible -i inventory all -m command -a "reboot"

After reboot, you can check the uptime

    ansible -i inventory all -m command -a "uptime"

Ensure a user "anna" is present

    ansible -i inventory all -m user -a "name=anna state=present"

Inventory is not needed as long as the inventory is specified in ansible.cfg

    ansible -m command -a "uptime"

## Exploring essential Ansible modules

Understanding module names

- Before Ansible 2.10, modules were referred to just by their name
  - `ansible all -m command -a reboot`
- Since Ansible 2.10, modules are a part of Ansible collections
  - Ansible all -m ansible.builtin.command -a reboot
- Collections allow for better management of large amounts of modules
- For compatibility reasons, the old module names can still be used

Listing modules

- pre-2.10: `ansible-doc -l`
- 2.10+: `ansible-doc -t module -l`

Exploring some essential modules

- `ansible.builtin.command` is the default module (no need to specify `-m`). Runs arbitraty commands on the managed nodes, but not trough a shell
- `ansible.builtin.shell` is the alternative to command, and should be used when shell items like pipes and redirects are required
- `ansible.builtin.raw` runs the command without the need for Python being present on the managed nodes
- `ansible.builtin.package` is a generic module for managing packages
- `ansible.builtin.user` allows for user management

- Always use the most specific module you can find
- Using generic modules often leads to problems as these don't do well in situations where the desired state already exists
- This is referred to as idempotency, meaning that if the current state already matches the desired state, no action is taken, and the module won't throw a fail.

```
ansible ansible29 -m command -a "rpm -qa | grep python" -> the pipe is ignored
ansible ansible29 -m shell -a "rpm -qa | grep python"
anbile all -m copy -a 'content="hello world" dest=/etc/motd'
  ansible all -a "cat /etc/motd"
ansible all -m package -a "name=nmap state=latest"
ansible rocky -m service -a "name=httpd state=started enabled=yes"
```

### Using module documentation

Using ansible-doc

- Module documentation is in `ansible-doc`
- Use `ansible-doc [-t module] -l` to show a list of all modules
- Use `ansible-doc modulename` for a detailed description of the module and its options
- The module's code can be found in the doc header
- Sections including an equals sign are mandatory for the module
- Notice the example section that exists at the end of the documentation for each of the modules
- See also section contains info about related modules

### Using Ansible in an indempotent way

Understanding idempotency

- Idempotence is a property of operations in mathematics and computer science whereby they can be applied multiple times without changing the result
- In Ansible it means that a module should run successfully, where it makes no difference if the current state already meets the desired state or not
- Many Ansible modules are idempotent by nature, but some are not (command, shell or raw)

```
# As user ansible on control, make sure you're in the directory that has inventory and ansible.cfg
ansible all -m ping
ansible all -m user -a "name=lisa"
ansible all -m command -a "useradd lisa"
ansible all -m user -a "name=lisa state=absent"
```

### Lab creating a user in Windows Server

```
# From windows folder...
ansible-doc -l | grep ^win
ansible-doc -l | grep ^win | grep facts
ansible win -m win_product_facts

ansible-doc win_user
    look for /mandatory
ansible win -m win_user -a "name=anna password=P@ssw0rd123!"
```

## 3 USING ANSIBLE PLAYBOOK