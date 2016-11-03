# go-angular

A little angular web frame created by Golang. It is easy to setup one-page Web site by Golang. It is a MVC web frame.   

# Why should you want this web frame?

- Jquery and Bootstrap and angular are integrated into the go-angular. But you can remove them from the go-angular if you donot want to use them.
- Go-angular is a light web frame for you to create a one-page web site. It have the easy-use and simple and powerful routes(controller).
- Full support angular. Server and client are separate.

# Usage:

## Dependencies:
- Golang http://golang.org/
- jQuery http://jquery.com/
- Angularjs https://angularjs.org/
- Twitter Bootstraphttp://twitter.github.com/bootstrap/

## Install and run:

    mkdir ~/go_codes/src
    cd ~/go_codes/src
    git clone https://github.com/DoBetter-pan/go-express.git
    cd go-angular
    source setenv.sh
    go build
    ./go-angular -host=127.0.0.1 -port=9898 
Then you can visit http://192.168.221.131:9898/

Note:
You must setup your golang environment first and will change your GOPATH when executing "source setenv.sh".
I tested it in Ubuntu. If you are using other OS, it is the same way as in Ubuntu. Please try.

## Develop:

I wrote a inline example: recipe. It is a sample in <AngularJS Up & Running>.
    
# What should we do next?
- Model will be supported. Go-angular will support sql and no-sql database as data source.
- More applications will be supported such as: blog, gallery etc.

# Support and contact

If you have any question and advice, you can cantact me in QQ Group: 536069420.
