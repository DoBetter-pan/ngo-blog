/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var app = angular.module('login', ['ngRoute', 'login.services', 'login.directives']);

app.config(['$interpolateProvider', function($interpolateProvider){
    $interpolateProvider.startSymbol('[[').endSymbol(']]');
}]);

app.config(['$routeProvider', function($routeProvider){
    $routeProvider.when('/view/:loginId', {
        controller: 'ViewCtrl',
        resolve: {
            login: function(LoginLoaderByCookie){
                return LoginLoaderByCookie();
            }
        },
        templateUrl: '/app/login/views/view.html'
    }).when('/new', {
        controller: 'NewCtrl',
        templateUrl: '/app/login/views/form.html'
    }).otherwise({redirectTo: '/view/100000'});
}]);
