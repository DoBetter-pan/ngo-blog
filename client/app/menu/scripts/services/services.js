/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var services = angular.module('menu.services', ['ngResource']);

/*get, save, query, remove, delete*/
services.factory('MenuSrv', ['$resource', function($resource){
    return $resource('/menusrv/:id', {id: '@id'});
}]);

services.factory('MultiMenuLoader', ['MenuSrv', '$q', function(MenuSrv, $q){
    return function() {
        var delay = $q.defer();
        MenuSrv.query(function(menus){
            delay.resolve(menus);
        }, function(){
            delay.reject('Unable to fetch menus');
        });
        return delay.promise;
    }
}]);

services.factory('MenuLoader', ['MenuSrv', '$route', '$q', function(MenuSrv, $route, $q){
    return function() {
        var delay = $q.defer();
        MenuSrv.get({id:$route.current.params.menuId}, function(menu){
            delay.resolve(menu);
        }, function(){
            delay.reject('Unable to fetch menu ' + $route.current.params.menuId);
        });
        return delay.promise;
    }
}]);

