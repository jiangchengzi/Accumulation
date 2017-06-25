#### calamari 开发环境的搭建

- 安装pyenv virtualenv pyenv-virtualenv

  ```shell
  $brew install pyenv
  $brew install pyenv-virtualenv
  $pip install virtualenv
  $eval "$(pyenv init -)" >> ~/.bash_profile
  $eval "$(pyenv virtualenv-init -)">> ~/.bash_profile
  $exec "$SHELL"
  $pyenv install 2.7.10
  $pyenv virtualenv 2.7.10 calamari
  $pyenv local calamari 
  #如果没有切换到下面这种格式，则需要运行pyenv actiate calamari一下：
  (calamari) houge-MacBook-Pro:calamari houge$ 
  ```

- 安装依赖

  ```shell
  $pip install -r requirements/2.7/requirements.txt
  ```

  ​

