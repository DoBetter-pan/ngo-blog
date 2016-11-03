/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var app = angular.module('category', ['ngRoute', 'category.services', 'category.directives']);

app.config(['$interpolateProvider', function($interpolateProvider){
    $interpolateProvider.startSymbol('[[').endSymbol(']]');
}]);

app.config(['$routeProvider', function($routeProvider){
    $routeProvider.when('/', {
        controller: 'ListCtrl',
        resolve: {
            categorys: function(MultiCategoryLoader) {
                return MultiCategoryLoader();
            }
        },
        templateUrl: '/app/category/views/list.html'
    }).when('/edit/:categoryId', {
        controller: 'EditCtrl',
    resolve: {
        category: function(CategoryLoader){
            return CategoryLoader();
        }
    },
    templateUrl: '/app/category/views/form.html'
    }).when('/view/:categoryId', {
        controller: 'ViewCtrl',
    resolve: {
        category: function(CategoryLoader) {
            return CategoryLoader();
        }
    },
    templateUrl: '/app/category/views/view.html'
    }).when('/new', {
        controller: 'NewCtrl',
    templateUrl: '/app/category/views/form.html'
    }).otherwise({redirectTo: '/'});
}]);
