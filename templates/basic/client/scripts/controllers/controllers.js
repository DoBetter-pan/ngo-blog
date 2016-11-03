/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

app.controller('ListCtrl', ['$scope', 'basics', function($scope, basics){
    //console.log(basics);
    $scope.basics = basics;
}]);

app.controller('ViewCtrl', ['$scope', '$location', 'basic', function($scope, $location, basic){
    $scope.basic = basic;

    $scope.edit = function(){
        $location.path('/edit/' + basic.id);
    }
}]);

app.controller('EditCtrl', ['$scope', '$location', 'basic', function($scope, $location, basic){
    $scope.basic = basic;

    $scope.save = function(){
        $scope.basic.$save(function(basic){
            $location.path('/view/' + basic.id);
        });
    };

    $scope.remove = function(){
        $scope.basic.$remove(function(basic){
            $location.path('/');
        });
    };
}]);

app.controller('NewCtrl', ['$scope', '$location', 'BasicSrv', function($scope, $location, BasicSrv){
    $scope.basic = new BasicSrv({
        id: -1
    });

    $scope.save = function(){
        $scope.basic.$save(function(basic){
            $location.path('/view/' + basic.id);
        });
    };
}]);

