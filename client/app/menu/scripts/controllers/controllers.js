/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

app.controller('ListCtrl', ['$scope', 'menus', function($scope, menus){
    //console.log(menus);
    $scope.menus = menus;
}]);

app.controller('ViewCtrl', ['$scope', '$location', 'menu', function($scope, $location, menu){
    $scope.menu = menu;

    $scope.edit = function(){
        $location.path('/edit/' + menu.id);
    }
}]);

app.controller('EditCtrl', ['$scope', '$location', 'menu', function($scope, $location, menu){
    $scope.menu = menu;

    $scope.save = function(){
        $scope.menu.$save(function(menu){
            $location.path('/view/' + menu.id);
        });
    };

    $scope.remove = function(){
        $scope.menu.$remove(function(menu){
            $location.path('/');
        });
    };
}]);

app.controller('NewCtrl', ['$scope', '$location', 'MenuSrv', function($scope, $location, MenuSrv){
    $scope.menu = new MenuSrv({
        id: -1
    });

    $scope.save = function(){
        $scope.menu.$save(function(menu){
            $location.path('/view/' + menu.id);
        });
    };
}]);

