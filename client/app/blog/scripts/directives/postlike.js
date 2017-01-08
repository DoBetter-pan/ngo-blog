/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

directives.directive('postlike', ['$http', function($http) {
    return {
        restrict: 'EA',
        replace: true,
        scope: {
            src: '@',
            pid: '@',                  
        }, 
        template: '<div class="postlike-action">' +
                        '<div class="postlike-position postlike-right">' + 
                           '<div class="postlike-like">' +
                                '<a class="postlike-likebg-style3 postlike-pointer" ng-click="like()">' +
                                    '<img src="/app/blog/styles/postlike/pixel.gif" title="顶">' +
                                        '<span class="postlike-likecount">{{ likecount }}</span>' +
                                '</a>' +
                            '</div>' +
                            '<div class="postlike-unlike">' +
                                '<a class="postlike-unlikebg-style3 postlike-pointer" ng-click="unlike()">' +
                                    '<img src="/app/blog/styles/postlike/pixel.gif" title="踩">' +
                                    '<span class="postlike-unlikecount">{{ unlikecount }}</span>' +
                                '</a>' +
                            '</div>' +
                        '</div>' +
                        '<div class="postlike-status postlike-right">{{ postlikestatus }}</div>' +
                    '</div>',
        controller: function($scope) {
            $scope.likecount = '+' + 0;
            $scope.unlikecount = '' - 0;
            $scope.postlikestatus = '  ';
              
            $http({method: 'GET', url: $scope.src, params:{pid:$scope.pid}}).then(function (result) {
                $scope.likecount = '+' + result.data.likecount; 
                $scope.unlikecount = '' + result.data.unlikecount;   
            }, function (result) {
                $scope.likecount = '+' + 0;
                $scope.unlikecount = '' - 0;
            });

            $scope.like = function() {
                $http({method: 'POST', url: $scope.src, params:{pid:$scope.pid, like:1}}).then(function (result) {
                    $scope.postlikestatus = result.data.message;    
                    $scope.likecount = '+' + result.data.likecount; 
                    $scope.unlikecount = '' + result.data.unlikecount;   
                }, function (result) {
                    $scope.postlikestatus = "Not Connected To Internet";                    
                    $scope.likecount = '+' + 0;
                    $scope.unlikecount = '' - 0;
                });
            };   

            $scope.unlike = function() {
              $http({method: 'POST', url: $scope.src, params:{pid:$scope.pid, unlike:1}}).then(function (result) {
                    $scope.postlikestatus = result.data.message;    
                    $scope.likecount = '+' + result.data.likecount; 
                    $scope.unlikecount = '' + result.data.unlikecount;   
                }, function (result) {
                    $scope.postlikestatus = "Not Connected To Internet";                    
                    $scope.likecount = '+' + 0;
                    $scope.unlikecount = '' - 0;
                });
            };         
        }
    }
}]);

