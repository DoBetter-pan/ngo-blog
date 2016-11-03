/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var services = angular.module('category.services', ['ngResource']);

/*get, save, query, remove, delete*/
services.factory('CategorySrv', ['$resource', function($resource){
    return $resource('/categorysrv/:id', {id: '@id'});
}]);

services.factory('MultiCategoryLoader', ['CategorySrv', '$q', function(CategorySrv, $q){
    return function() {
        var delay = $q.defer();
        CategorySrv.query(function(categorys){
            delay.resolve(categorys);
        }, function(){
            delay.reject('Unable to fetch categorys');
        });
        return delay.promise;
    }
}]);

services.factory('CategoryLoader', ['CategorySrv', '$route', '$q', function(CategorySrv, $route, $q){
    return function() {
        var delay = $q.defer();
        CategorySrv.get({id:$route.current.params.categoryId}, function(category){
            delay.resolve(category);
        }, function(){
            delay.reject('Unable to fetch category ' + $route.current.params.categoryId);
        });
        return delay.promise;
    }
}]);

