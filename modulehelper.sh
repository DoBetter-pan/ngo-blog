#!/bin/bash

# author: yingx
# date: 2016-08-18

function usage() {
    echo "USAGE: $0 create|remove modulename templatename"
    exit 1
}

#function capitalize() {
#    split=( $1 )
#    echo "${split[*]^}"
#}

function capitalize() {
    echo $1 | sed 's/./\U&/'
}

function substr() {
    local str1=$1
    local str2=$2

    if [[ ${str1/${str2}//} == $str1 ]]; then
        return 0
    else
        return 1
    fi
}

function checkParams() {
    case "$1" in
        "create")
            if [ ! -d $5/$2 ]; then
                echo "The template:$2 does not exist!"
                echo ""
                usage
            fi
            if [ -f $4/${3}Controller.go ]; then
                echo "The module:$3 has already existed, please try another name!"
                echo ""
                usage
            fi
            ;;
        "remove")
            ;;
        *)
            echo "The operation:$1 is not supported!"
            echo ""
            usage
            ;;
    esac
}

function createModule() {
    echo "USAGE: $0 create|remove templatename modulename"
}

function removeModule() {
    echo "USAGE: $0 create|remove templatename modulename"
}

#retrieve params
case "$#" in
    "2")
        operation=$1
        modulename=$2
        templatename=basic
        ;;
    "3")
        operation=$1
        modulename=$2
        templatename=$3
        ;;
    *)
        echo "Parameter count is not right!"
        echo ""
        usage
        ;;
esac

basedir=`pwd`
controllerpath=$basedir/server/controller
modelpath=$basedir/server/model
datapath=$basedir/server/data

templatespath=$basedir/templates
templatescontrollerpath=$templatespath/$templatename/server/controller
templatesmodelpath=$templatespath/$templatename/server/model
templatesdatapath=$templatespath/$templatename/server/data

#check params
checkParams $operation $templatename $modulename $controllerpath $templatespath

case "$operation" in
    "create")
        echo "Creating a new module $modulename by template $templatename"
        templatenameCap=`capitalize $templatename`
        modulenameCap=`capitalize $modulename`

        #server
        moduleController=$controllerpath/${modulename}Controller.go
        echo "Cteate $moduleController ... "
        cp $templatescontrollerpath/${templatename}Controller.go $moduleController
        sed -i "s/$templatename/$modulename/g" $moduleController
        sed -i "s/$templatenameCap/$modulenameCap/g" $moduleController

        moduleSrvController=$controllerpath/${modulename}SrvController.go
        echo "Cteate $moduleSrvController ... "
        cp $templatescontrollerpath/${templatename}SrvController.go $moduleSrvController
        sed -i "s/$templatename/$modulename/g" $moduleSrvController
        sed -i "s/$templatenameCap/$modulenameCap/g" $moduleSrvController

        index=0
        sqlflag=0
        createsqlitems=""
        selectsqlitems=""
        insertsqlcolumn=""
        insertsqlvalue=""
        updatesqlitems=""
        jsonitems="    "
        scanitemlist=""
        insertitemlist=""
        templateSrvData=$templatesdatapath/${templatename}SrvData.txt
        while IFS='|' read -ra sqlcolumn; do
            if [ $index -ne 0 ]; then
                createsqlitems=${createsqlitems},
                selectsqlitems=${selectsqlitems},
                jsonitems="${jsonitems}\n    "
                scanitemlist=${scanitemlist},
            fi
            for info in "${sqlcolumn[@]}"; do
                createsqlitems="${createsqlitems} ${info}"
            done
            columnNameCap=`capitalize ${sqlcolumn[0]}`
            if [ $sqlflag -ne 0 ]; then
                insertsqlcolumn=${insertsqlcolumn},
                insertsqlvalue=${insertsqlvalue},
                updatesqlitems=${updatesqlitems},
                insertitemlist=${insertitemlist},
            fi
            selectsqlitems="${selectsqlitems} ${sqlcolumn[0]}"
            scanitemlist="${scanitemlist} \&${modulename}.${columnNameCap}"
            if [ ${sqlcolumn[0]} != "id" ]; then
                insertsqlcolumn="${insertsqlcolumn} ${sqlcolumn[0]}"
                insertsqlvalue="${insertsqlvalue} ?"
                updatesqlitems="${updatesqlitems} ${sqlcolumn[0]}=?"
                insertitemlist="${insertitemlist} ${modulename}.${columnNameCap}"
                sqlflag=1
            fi
            substr ${sqlcolumn[1]} "char"
            if [ $? -eq 1 ]; then
                jsonitems="${jsonitems} ${columnNameCap}  string \`json:\"${sqlcolumn[0]}\"\`"
            else
                jsonitems="${jsonitems} ${columnNameCap} int64 \`json:\"${sqlcolumn[0]}\"\`"
            fi
            index=$(( $index + 1 ))
        done < $templateSrvData
        insertitemlist="${insertitemlist}, ${modulename}.Id"
        
        createsql="create table ${modulename} ($createsqlitems)"
        moduleSrvData=$datapath/${modulename}SrvData.sql
        echo "Cteate $moduleSrvData ... "
        echo $createsql > $moduleSrvData

        moduleSrvModel=$modelpath/${modulename}SrvModel.go
        echo "Cteate $moduleSrvModel ... "
        cp $templatesmodelpath/${templatename}SrvModel.go $moduleSrvModel
        sed -i "s/$templatename/$modulename/g" $moduleSrvModel
        sed -i "s/$templatenameCap/$modulenameCap/g" $moduleSrvModel
        sed -i "s/##Item##/${jsonitems}/g" $moduleSrvModel
        selectsql="select${selectsqlitems} from ${modulename}"
        selectonesql="select${selectsqlitems} from ${modulename} where id=?"
        insertsql="insert into ${modulename}(${insertsqlcolumn}) values(${insertsqlvalue})"
        updatesql="update ${modulename} set${updatesqlitems} where id=?"
        deletesql="delete from ${modulename} where id=?"
        sed -i "s/##query##/${selectsql}/g" $moduleSrvModel
        sed -i "s/##queryone##/${selectonesql}/g" $moduleSrvModel
        sed -i "s/##insert##/${insertsql}/g" $moduleSrvModel
        sed -i "s/##update##/${updatesql}/g" $moduleSrvModel
        sed -i "s/##delete##/${deletesql}/g" $moduleSrvModel
        sed -i "s/##scanitemlist##/${scanitemlist}/g" $moduleSrvModel
        sed -i "s/##insertitemlist##/${insertitemlist}/g" $moduleSrvModel
        sed -i "s/##updateitemlist##/${insertitemlist}/g" $moduleSrvModel


        moduleHandler="func ${modulename}Handler(w http.ResponseWriter, r *http.Request) {\n"
        moduleHandler="${moduleHandler}    ${modulename} := controller.New${modulenameCap}Controller()\n"
        moduleHandler="${moduleHandler}    controller := reflect.ValueOf(${modulename})\n"
        moduleHandler="${moduleHandler}    controllerAction(w, r, func() reflect.Value {\n"
        moduleHandler="${moduleHandler}        return controller\n"
        moduleHandler="${moduleHandler}        })\n"
        moduleHandler="${moduleHandler}}\n\n"

        moduleSrvHandler="func ${modulename}SrvHandler(w http.ResponseWriter, r *http.Request) {\n"
        moduleSrvHandler="${moduleSrvHandler}    ${modulename}Srv := controller.New${modulenameCap}SrvController()\n"
        moduleSrvHandler="${moduleSrvHandler}    controller := reflect.ValueOf(${modulename}Srv)\n"
        moduleSrvHandler="${moduleSrvHandler}    controllerResty(w, r, func() reflect.Value {\n"
        moduleSrvHandler="${moduleSrvHandler}        return controller\n"
        moduleSrvHandler="${moduleSrvHandler}        })\n"
        moduleSrvHandler="${moduleSrvHandler}}\n\n"

        sed -i "/func main/s/^/$moduleHandler/g" $basedir/app.go
        sed -i "/func main/s/^/$moduleSrvHandler/g" $basedir/app.go

        moduleHandlerPath="    http.HandleFunc\(\"\/${modulename}\/\", ${modulename}Handler\)\n"
        moduleSrvHandlerPath="    http.HandleFunc\(\"\/${modulename}srv\/\", ${modulename}SrvHandler\)\n"
        sed -i "/server := fmt.Sprintf/s/^/${moduleHandlerPath}/g" $basedir/app.go
        sed -i "/server := fmt.Sprintf/s/^/${moduleSrvHandlerPath}/g" $basedir/app.go

        #client
        clientappmodule=${basedir}/client/app/${modulename}
        cp ${templatespath}/${templatename}/client ${clientappmodule} -fr
        sed -i "s/$templatename/$modulename/g" ${clientappmodule}/index.html
        sed -i "s/$templatenameCap/$modulenameCap/g" ${clientappmodule}/index.html
        sed -i "s/$templatename/$modulename/g" ${clientappmodule}/views/list.html
        sed -i "s/$templatenameCap/$modulenameCap/g" ${clientappmodule}/views/list.html
        sed -i "s/$templatename/$modulename/g" ${clientappmodule}/views/view.html
        sed -i "s/$templatenameCap/$modulenameCap/g" ${clientappmodule}/views/view.html
        sed -i "s/$templatename/$modulename/g" ${clientappmodule}/views/form.html
        sed -i "s/$templatenameCap/$modulenameCap/g" ${clientappmodule}/views/form.html
        sed -i "s/$templatename/$modulename/g" ${clientappmodule}/scripts/app.js
        sed -i "s/$templatenameCap/$modulenameCap/g" ${clientappmodule}/scripts/app.js
        sed -i "s/$templatename/$modulename/g" ${clientappmodule}/scripts/services/services.js
        sed -i "s/$templatenameCap/$modulenameCap/g" ${clientappmodule}/scripts/services/services.js
        sed -i "s/$templatename/$modulename/g" ${clientappmodule}/scripts/directives/directives.js
        sed -i "s/$templatenameCap/$modulenameCap/g" ${clientappmodule}/scripts/directives/directives.js
        sed -i "s/$templatename/$modulename/g" ${clientappmodule}/scripts/controllers/controllers.js
        sed -i "s/$templatenameCap/$modulenameCap/g" ${clientappmodule}/scripts/controllers/controllers.js

        ;;
    "remove")
        echo "Remove the existed module $modulename"
        #server
        rm $controllerpath/${modulename}Controller.go
        rm $controllerpath/${modulename}SrvController.go
        rm $modelpath/${modulename}SrvModel.go
        rm $datapath/${modulename}SrvData.sql

        moduleHandler="func ${modulename}Handler"
        moduleSrvHandler="func ${modulename}SrvHandler"

        sed -i "/${moduleHandler}/{N;N;N;N;N;N;N;d}" $basedir/app.go
        sed -i "/${moduleSrvHandler}/{N;N;N;N;N;N;N;d}" $basedir/app.go

        moduleHandlerPath="http.HandleFunc(\"\/${modulename}\/\", ${modulename}Handler)"
        moduleSrvHandlerPath="http.HandleFunc(\"\/${modulename}srv\/\", ${modulename}SrvHandler)"

        sed -i "/${moduleHandlerPath}/d" $basedir/app.go
        sed -i "/${moduleSrvHandlerPath}/d" $basedir/app.go

        #client
        rm ${basedir}/client/app/${modulename} -fr


        ;;
esac



