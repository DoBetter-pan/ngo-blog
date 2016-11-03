/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var services = angular.module('login.services', ['ngResource']);

/*get, save, query, remove, delete*/
services.factory('LoginSrv', ['$resource', function($resource){
    return $resource('/loginsrv/:id', {id: '@id'}, {
        checkUser: {method: 'POST', url: '/loginsrv/checkuser', params: {p: 1}, isArray: false},
        getUserByCookie: {method: 'GET', url: '/loginsrv/getuser', params: {p: 1}, isArray: false}
    });
}]);

services.factory('MultiLoginLoader', ['LoginSrv', '$q', function(LoginSrv, $q){
    return function() {
        var delay = $q.defer();
        LoginSrv.query(function(logins){
            delay.resolve(logins);
        }, function(){
            delay.reject('Unable to fetch logins');
        });
        return delay.promise;
    }
}]);

services.factory('LoginLoader', ['LoginSrv', '$route', '$q', function(LoginSrv, $route, $q){
    return function() {
        var delay = $q.defer();
        LoginSrv.get({id:$route.current.params.loginId}, function(login){
            delay.resolve(login);
        }, function(){
            delay.reject('Unable to fetch login ' + $route.current.params.loginId);
        });
        return delay.promise;
    }
}]);


services.factory('LoginLoaderByCookie', ['LoginSrv', '$route', '$q', function(LoginSrv, $route, $q){
    return function() {
        var delay = $q.defer();
        LoginSrv.getUserByCookie({id:$route.current.params.loginId}, function(login){
            delay.resolve(login);
        }, function(){
            delay.reject('Unable to fetch login ' + $route.current.params.loginId);
        });
        return delay.promise;
    }
}]);
