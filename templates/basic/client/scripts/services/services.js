/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var services = angular.module('basic.services', ['ngResource']);

/*get, save, query, remove, delete*/
services.factory('BasicSrv', ['$resource', function($resource){
    return $resource('/basicsrv/:id', {id: '@id'});
}]);

services.factory('MultiBasicLoader', ['BasicSrv', '$q', function(BasicSrv, $q){
    return function() {
        var delay = $q.defer();
        BasicSrv.query(function(basics){
            delay.resolve(basics);
        }, function(){
            delay.reject('Unable to fetch basics');
        });
        return delay.promise;
    }
}]);

services.factory('BasicLoader', ['BasicSrv', '$route', '$q', function(BasicSrv, $route, $q){
    return function() {
        var delay = $q.defer();
        BasicSrv.get({id:$route.current.params.basicId}, function(basic){
            delay.resolve(basic);
        }, function(){
            delay.reject('Unable to fetch basic ' + $route.current.params.basicId);
        });
        return delay.promise;
    }
}]);

