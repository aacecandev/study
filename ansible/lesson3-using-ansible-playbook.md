## 3 USING ANSIBLE PLAYBOOK

### 3.1 Understanding playbook structure

Understanding playbooks

- Playbooks are used to run multiple tasks in order on multiple hosts
- Playbooks are the way to go to define dependency relations between tasks
- Playbooks can also be used to implement conditional structures, where tasks are only executed if specific conditions are met

Understanding playbook elements

- In a playbook, one or more plays are defined
- Each play has a header that defines different properties, including the hosts on which it has to run
- In a play, one or more tasks are defined
- Playbooks are written in YAML format (doesn't support tabs for indentation)

```
# configure vim for indentation
# vim ~/.vimrc

autocmd FileType yaml setlocal ai ts=2 sw=2 expandtab
```

- The relations are defined using indentation
- Task properties are defined using key: value pairs
- Most keys can have one string value
- Some keys have a list value. Then, each item is identified using an hyphen


```yaml
- name: myplay
  hosts: rocky
  tasks:
  - name: task1
    debug:
      msg: "Hello world 2" 
  - name: task2
    debug:
      msg: "Hello world 1"

- name: "ubuntu settings"
  hosts: ubuntu
  tasks:
  - name: hello ubuntu
    debug:
      msg: "Hello world ubuntu"
```