/** 
* @file filters.js
* @brief filters
* @author yingx
* @date 2016-11-09
*/

var filters = angular.module('util.filters', []);

filters.filter('rawHtml', ['$sce', function($sce){
    return function(val) {
        return $sce.trustAsHtml(val);
    };
}]);
