/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var app = angular.module('basic', ['ngRoute', 'basic.services', 'basic.directives']);

app.config(['$interpolateProvider', function($interpolateProvider){
    $interpolateProvider.startSymbol('[[').endSymbol(']]');
}]);

app.config(['$routeProvider', function($routeProvider){
    $routeProvider.when('/', {
        controller: 'ListCtrl',
        resolve: {
            basics: function(MultiBasicLoader) {
                return MultiBasicLoader();
            }
        },
        templateUrl: '/app/basic/views/list.html'
    }).when('/edit/:basicId', {
        controller: 'EditCtrl',
    resolve: {
        basic: function(BasicLoader){
            return BasicLoader();
        }
    },
    templateUrl: '/app/basic/views/form.html'
    }).when('/view/:basicId', {
        controller: 'ViewCtrl',
    resolve: {
        basic: function(BasicLoader) {
            return BasicLoader();
        }
    },
    templateUrl: '/app/basic/views/view.html'
    }).when('/new', {
        controller: 'NewCtrl',
    templateUrl: '/app/basic/views/form.html'
    }).otherwise({redirectTo: '/'});
}]);
