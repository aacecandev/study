## libselinux-python3


Ansible playbook fails with libselinux-python aren't installed on RHEL8
Solution Verified - Updated December 30 2020 at 5:47 PM - English
Environment
  Red Hat Enterprise Linux 8
  Ansible Engine 2.8+

Issue
    When running an Ansible playbook against an RHEL8 system, it throws the error below even after the python3-libselinux. package has been installed

```
Aborting, target uses selinux but python bindings (libselinux-python) aren't installed!".
```

Resolution
  Verify if the package python3-libselinux is installed and which python version it is linked to:

```
# rpm -q python3-libselinux
# rpm -ql python3-libselinux.x86_64
/usr/lib/.build-id
/usr/lib/.build-id/08
/usr/lib/.build-id/08/472d4d1da85214c069b4c90b5bede2cf044de3
/usr/lib/.build-id/e4
/usr/lib/.build-id/e4/0257977daf7236760a7e4747adf60979aecfd7
/usr/lib64/python3.6/site-packages/_selinux.cpython-36m-x86_64-linux-gnu.so
/usr/lib64/python3.6/site-packages/selinux
/usr/lib64/python3.6/site-packages/selinux-2.9-py3.6.egg-info
/usr/lib64/python3.6/site-packages/selinux/__init__.py
/usr/lib64/python3.6/site-packages/selinux/__pycache__
/usr/lib64/python3.6/site-packages/selinux/__pycache__/__init__.cpython-36.opt-1.pyc
/usr/lib64/python3.6/site-packages/selinux/__pycache__/__init__.cpython-36.pyc
/usr/lib64/python3.6/site-packages/selinux/audit2why.cpython-36m-x86_64-linux-gnu.so
```

Modify the alternatives configuration to point to the default Python version covered by the python3-libselinux package.

```
# python3 --version
Python 3.8.3

# alternatives  --list  | grep python
python                  manual  /usr/libexec/no-python
python3                 auto    /usr/bin/python3.8                 <=== must match with the python3-libselinux package

# alternatives  --config python

There are 3 programs which provide 'python'.

  Selection    Command
-----------------------------------------------
*+ 1           /usr/libexec/no-python
   2           /usr/bin/python3                       <=== select this one
   3           /usr/bin/python3.8

    Confirm Python has the correct version

# python --version
Python 3.6.8
```

Root Cause

The current Python version installed on RHEL8 is Python 3.6. Whenever this error happens, most likely the default python version on the managed node has been modified to a different version which is not covered by the default python3-libselinux RPM package which is intended for Python 3.6.

Diagnostic Steps

Verify if the system has multiple Python installed

```
# rpm -aq | grep python3?-3.
python38-3.8.3-3.module+el8.3.0+7680+79e7e61a.x86_64  
python36-3.6.8-2.module+el8.1.0+3334+5cb623d7.x86_64 
```

Check the alternatives configuration to verify which is the default Python version the system is pointing to:

```
# alternatives  --config python
* 
There are 3 programs which provide 'python'.

  Selection    Command
-----------------------------------------------
*+ 1           /usr/libexec/no-python
   2           /usr/bin/python3
   3           /usr/bin/python3.8

# ls -la /etc/alternatives/python*
lrwxrwxrwx. 1 root root 18 Dec 29 17:47 /etc/alternatives/python -> /usr/bin/python3.8
lrwxrwxrwx. 1 root root 18 Dec 29 17:43 /etc/alternatives/python3 -> /usr/bin/python3.6
lrwxrwxrwx. 1 root root 34 Dec 29 17:43 /etc/alternatives/python3-man -> /usr/share/man/man1/python3.6.1.gz
```

