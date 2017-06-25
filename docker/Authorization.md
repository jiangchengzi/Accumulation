Authorization

reg-





pyenv无法activate的原因是pyenv-virtualenv没有初始化

```shell
$ echo 'eval "$(pyenv virtualenv-init -)"' >> ~/.bash_profile
$ exec "$SHELL"
```



