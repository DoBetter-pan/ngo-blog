/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var services = angular.module('section.services', ['ngResource']);

/*get, save, query, remove, delete*/
services.factory('SectionSrv', ['$resource', function($resource){
    return $resource('/sectionsrv/:id', {id: '@id'});
}]);

services.factory('MultiSectionLoader', ['SectionSrv', '$q', function(SectionSrv, $q){
    return function() {
        var delay = $q.defer();
        SectionSrv.query(function(sections){
            delay.resolve(sections);
        }, function(){
            delay.reject('Unable to fetch sections');
        });
        return delay.promise;
    }
}]);

services.factory('SectionLoader', ['SectionSrv', '$route', '$q', function(SectionSrv, $route, $q){
    return function() {
        var delay = $q.defer();
        SectionSrv.get({id:$route.current.params.sectionId}, function(section){
            delay.resolve(section);
        }, function(){
            delay.reject('Unable to fetch section ' + $route.current.params.sectionId);
        });
        return delay.promise;
    }
}]);

