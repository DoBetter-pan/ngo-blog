/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var admin = angular.module('admin', ['ngRoute', 'blog.services', 'section.services', 'category.services', 'util.directives' ]);

admin.config(['$interpolateProvider', function($interpolateProvider){
    $interpolateProvider.startSymbol('[[').endSymbol(']]');
}]);

admin.config(['$routeProvider', function($routeProvider){
    $routeProvider.when('/newblog', {
        controller: 'NewBlogCtrl',
        resolve: {
            sections: function(MultiSectionLoader){
                return MultiSectionLoader();
            },
            categories: function(MultiCategoryLoader) {
                return MultiCategoryLoader();
            }
        },
        templateUrl: '/app/blog/views/newblog.html'
    }).when('/listblog', {
        controller: 'ListCtrl',
        resolve: {
            articles: function(MultiBlogLoader) {
                return MultiBlogLoader();
            }
        },
        templateUrl: '/app/blog/views/list.html'
    }).when('/editblog/:blogId', {
        controller: 'EditCtrl',
        resolve: {
            article: function(BlogLoader){
                return BlogLoader();
            }
        },
        templateUrl: '/app/blog/views/form.html'
    }).when('/viewblog/:blogId', {
        controller: 'ViewBlogCtrl',
        resolve: {
            article: function(BlogLoader) {
                return BlogLoader();
            }
        },
        templateUrl: '/app/blog/views/viewblog.html'
    }).otherwise({redirectTo: '/newblog'});
}]);
