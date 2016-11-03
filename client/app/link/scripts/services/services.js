/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var services = angular.module('link.services', ['ngResource']);

/*get, save, query, remove, delete*/
services.factory('LinkSrv', ['$resource', function($resource){
    return $resource('/linksrv/:id', {id: '@id'});
}]);

services.factory('MultiLinkLoader', ['LinkSrv', '$q', function(LinkSrv, $q){
    return function() {
        var delay = $q.defer();
        LinkSrv.query(function(links){
            delay.resolve(links);
        }, function(){
            delay.reject('Unable to fetch links');
        });
        return delay.promise;
    }
}]);

services.factory('LinkLoader', ['LinkSrv', '$route', '$q', function(LinkSrv, $route, $q){
    return function() {
        var delay = $q.defer();
        LinkSrv.get({id:$route.current.params.linkId}, function(link){
            delay.resolve(link);
        }, function(){
            delay.reject('Unable to fetch link ' + $route.current.params.linkId);
        });
        return delay.promise;
    }
}]);

