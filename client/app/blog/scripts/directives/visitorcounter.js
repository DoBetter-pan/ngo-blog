/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

directives.directive('visitorCounter', ['$http', function($http) {
    return {
        restrict: 'EA',
        replace: true,
        scope: {
            src: '@',                  
        }, 
        template: '<div class="counter">访问人数：{{ counter }}</div>',
        controller: function($scope) {
            $scope.counter = 0;
              
            $http({method: 'GET', url: $scope.src}).then(function (result) {
                $scope.counter = result.data.visitorCounter;                          
            }, function (result) {
                $scope.counter = 0;
            });            
        }
    }
}]);