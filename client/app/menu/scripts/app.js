/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var app = angular.module('menu', ['ngRoute', 'menu.services', 'menu.directives']);

app.config(['$interpolateProvider', function($interpolateProvider){
    $interpolateProvider.startSymbol('[[').endSymbol(']]');
}]);

app.config(['$routeProvider', function($routeProvider){
    $routeProvider.when('/', {
        controller: 'ListCtrl',
        resolve: {
            menus: function(MultiMenuLoader) {
                return MultiMenuLoader();
            }
        },
        templateUrl: '/app/menu/views/list.html'
    }).when('/edit/:menuId', {
        controller: 'EditCtrl',
    resolve: {
        menu: function(MenuLoader){
            return MenuLoader();
        }
    },
    templateUrl: '/app/menu/views/form.html'
    }).when('/view/:menuId', {
        controller: 'ViewCtrl',
    resolve: {
        menu: function(MenuLoader) {
            return MenuLoader();
        }
    },
    templateUrl: '/app/menu/views/view.html'
    }).when('/new', {
        controller: 'NewCtrl',
    templateUrl: '/app/menu/views/form.html'
    }).otherwise({redirectTo: '/'});
}]);
