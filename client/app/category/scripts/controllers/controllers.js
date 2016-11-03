/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

app.controller('ListCtrl', ['$scope', 'categorys', function($scope, categorys){
    //console.log(categorys);
    $scope.categorys = categorys;
}]);

app.controller('ViewCtrl', ['$scope', '$location', 'category', function($scope, $location, category){
    $scope.category = category;

    $scope.edit = function(){
        $location.path('/edit/' + category.id);
    }
}]);

app.controller('EditCtrl', ['$scope', '$location', 'category', function($scope, $location, category){
    $scope.category = category;

    $scope.save = function(){
        $scope.category.$save(function(category){
            $location.path('/view/' + category.id);
        });
    };

    $scope.remove = function(){
        $scope.category.$remove(function(category){
            $location.path('/');
        });
    };
}]);

app.controller('NewCtrl', ['$scope', '$location', 'CategorySrv', function($scope, $location, CategorySrv){
    $scope.category = new CategorySrv({
        id: -1
    });

    $scope.save = function(){
        $scope.category.$save(function(category){
            $location.path('/view/' + category.id);
        });
    };
}]);

