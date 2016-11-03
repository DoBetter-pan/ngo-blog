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

blog.controller('ListCtrl', ['$scope', 'articles', function($scope, articles){
    //console.log(articles);
    $scope.articles = articles;
}]);

blog.controller('ViewCtrl', ['$scope', 'article', function($scope, article){
    $scope.article = article;
}]);
