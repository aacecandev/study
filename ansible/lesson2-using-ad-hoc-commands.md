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

## Examples

This will fail as it will reboot and then request status

    ansible -i inventory all -m command -a "reboot"

After reboot, you can check the uptime

    ansible -i inventory all -m command -a "uptime"

Ensure a user "anna" is present

    ansible -i inventory all -m user -a "name=anna state=present"

Inventory is not needed as long as the inventory is specified in ansible.cfg

    ansible -m command -a "uptime"

