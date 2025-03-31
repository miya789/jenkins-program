# Project Title

A brief description of what this project does and who it's for.

## Features

- Feature 1
- Feature 2
- Feature 3

## Installation

1. Clone the repository:
  ```bash
  git clone https://github.com/your-username/your-repo.git
  ```
2. Navigate to the project directory:
  ```bash
  cd your-repo
  ```
3. Install dependencies:
  ```bash
  npm install
  ```

## Usage

Provide instructions and examples for using the project.

1. 対象Jenkinsの用意
  ```bash
  ./jenkins-run.sh
  ```

2. 初期設定
  デフォルトのプラグインを全てインストール

3. 行列の権限を設定

4. adminのアクセストークンを発行

5. コマンドの実行
  ```bash
  go build main.go
  ```

6. ホストで実行
  ```bash
  export JENKINS_USER=xxx
  export JENKINS_TOKEN=xxx
  ./main
  ```
