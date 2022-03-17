## 그동안 공부했던 golang 레포지토리를 합쳤습니다.
## Github 예전 내 레포지토리들을 한 레포지토리로 합치는 방법
    1) 기본 레포지토리들을 합칠 새로운 레포지토리를 만듭니다.
    2) 로컬에서 새로운 폴더를 만들고 새로 만든 레포지토리를 git clone 해줍니다.
    ex> git clone https://github.com/블라블라
    3) cd로 .git이 존재하는 해당 레포지토리 폴더로 이동합니다.
    4) README.md등 파일을 만들어 아무런 커밋이나 우선 진행합니다. 선 커밋을 안할 시 unknown revision or path not in the working tree라는 에러 문구를 만날 수 있습니다.
    4) git subtree add --prefix=(새로만든레포지토리 하위폴더로 생성할 이름) (옮겨서 합쳐질 예전 레포지토리 주소) branch명 을 입력해서 합칩니다.
    ex> git subtree add --prefix=Baekjoon https://github.com/rapael3020/answer.git main
### 레포지토리
- gotuto1 폴더에는 2020.05.29에 생성된 폴더로, golang을 이용한 file서버 및 테스트 파일의 사용으로 예제 실습을 진행하였습니다.
- gotuto2 폴더에는 2020.05.12에 생성된 폴더로, golang과 Docker 실습 파일입니다.
- gotuto3 폴더에는 2020.05.16에 생성된 폴더로, golang을 활용한 pingpong서버 예제 실습 파일입니다.
- gotuto4 폴더에는 2020.07.15에 생성된 폴더로, golang을 활용한 web개발 예제 실습 파일입니다.