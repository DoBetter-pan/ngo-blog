/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var services = angular.module('blog.services', ['ngResource']);

/*get, save, query, remove, delete*/
services.factory('BlogSrv', ['$resource', function($resource){
    return $resource('/blogsrv/:id', {id: '@id'});
}]);

services.factory('MultiBlogLoader', ['BlogSrv', '$route', '$q', function(BlogSrv, $route, $q){
    return function() {
        var delay = $q.defer();
        BlogSrv.query({s:$route.current.params.s, c:$route.current.params.c, id:$route.current.params.id}, function(articles){
            delay.resolve(articles);
        }, function(){
            delay.reject('Unable to fetch articles');
        });
        return delay.promise;
    }
}]);

services.factory('BlogLoader', ['BlogSrv', '$route', '$q', function(BlogSrv, $route, $q){
    return function() {
        var delay = $q.defer();
        BlogSrv.get({id:$route.current.params.blogId}, function(blog){
            delay.resolve(blog);
        }, function(){
            delay.reject('Unable to fetch blog ' + $route.current.params.blogId);
        });
        return delay.promise;
    }
}]);

