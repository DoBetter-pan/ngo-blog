/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

app.controller('ViewCtrl', ['$scope', '$rootScope', '$location', 'login', function($scope, $rootScope, $location, login){
    $scope.login = login;

    var changeLocation = function(url, forceReload) {
        $scope = $scope || angular.element(document).scope();
        if(forceReload || $scope.$$phase) {
            window.location = url;
        }
        else {
            //only use this if you want to replace the history stack
            //$location.path(url).replace();

            //this this if you want to change the URL and add it to the history stack
            $location.path(url);
            $scope.$apply();
        }
    };

    $scope.checkUser = function(){
        $scope.login.$checkUser(function(login){
            if(login.status == 0) {
                //$location.url('/blog');
                changeLocation('/blog/admin', true);
            }
        });
    }
}]);

app.controller('NewCtrl', ['$scope', '$location', 'LoginSrv', function($scope, $location, LoginSrv){
    $scope.login = new LoginSrv({
        id: -1
    });

    $scope.save = function(){
        $scope.login.$save(function(login){
            $location.path('/view/' + login.id);
        });
    };
}]);

