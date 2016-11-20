/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

blog.controller('IndexCtrl', ['$scope', 'articles', function($scope, articles){
    //console.log(articles);
    $scope.articles = articles;
}]);

blog.controller('ListCtrl', ['$scope', '$timeout', 'articles', function($scope, $timeout, articles){
    //console.log(articles);
    $scope.articles = articles;
    $scope.$on('$viewContentLoaded', function(){
        $timeout(function() {
            $("pre").addClass("prettyprint");
            prettyPrint(); 
        }, 0);
    });
}]);

blog.controller('ViewCtrl', ['$scope', '$timeout', 'article', function($scope, $timeout, article){
    $scope.article = article;
    $scope.$on('$viewContentLoaded', function(){
        $timeout(function() {
            $("pre").addClass("prettyprint");
            prettyPrint(); 
        }, 0);
    });
}]);
