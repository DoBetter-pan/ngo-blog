/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var blog = angular.module('blog', ['ngRoute', 'blog.services', 'util.directives']);

blog.config(['$interpolateProvider', function($interpolateProvider){
    $interpolateProvider.startSymbol('[[').endSymbol(']]');
}]);

blog.config(['$routeProvider', function($routeProvider){
    $routeProvider.when('/', {
        controller: 'IndexCtrl',
        resolve: {
            articles: function(MultiBlogLoader) {
                return MultiBlogLoader();
            }
        },
        templateUrl: '/app/blog/views/index.html'
    }).when('/list', {
        controller: 'ListCtrl',
        resolve: {
            articles: function(MultiBlogLoader) {
                return MultiBlogLoader();
            }
        },
        templateUrl: '/app/blog/views/list.html'
    }).when('/view/:blogId', {
        controller: 'ViewCtrl',
        resolve: {
            article: function(BlogLoader) {
                return BlogLoader();
            }
        },
        templateUrl: '/app/blog/views/view.html'
    }).otherwise({redirectTo: '/'});
}]);
