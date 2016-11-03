/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var app = angular.module('link', ['ngRoute', 'link.services', 'link.directives']);

app.config(['$interpolateProvider', function($interpolateProvider){
    $interpolateProvider.startSymbol('[[').endSymbol(']]');
}]);

app.config(['$routeProvider', function($routeProvider){
    $routeProvider.when('/', {
        controller: 'ListCtrl',
        resolve: {
            links: function(MultiLinkLoader) {
                return MultiLinkLoader();
            }
        },
        templateUrl: '/app/link/views/list.html'
    }).when('/edit/:linkId', {
        controller: 'EditCtrl',
    resolve: {
        link: function(LinkLoader){
            return LinkLoader();
        }
    },
    templateUrl: '/app/link/views/form.html'
    }).when('/view/:linkId', {
        controller: 'ViewCtrl',
    resolve: {
        link: function(LinkLoader) {
            return LinkLoader();
        }
    },
    templateUrl: '/app/link/views/view.html'
    }).when('/new', {
        controller: 'NewCtrl',
    templateUrl: '/app/link/views/form.html'
    }).otherwise({redirectTo: '/'});
}]);
