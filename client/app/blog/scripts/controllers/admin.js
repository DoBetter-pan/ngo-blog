/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

admin.controller('NewBlogCtrl', ['$scope', '$location', 'BlogSrv', 'sections', 'categories', function($scope, $location, BlogSrv, sections, categories){
    $scope.sections = sections;
    $scope.categories = categories;
    $scope.categoriesBySec = categories.slice(0);
    $scope.changeSeciton = function(secId) {
        $scope.categoriesBySec.splice(0, $scope.categoriesBySec.length);
        $scope.categories.forEach(function(e){
            if(e.sectionId == secId){
                $scope.categoriesBySec.push(e);
            }
        });
        $scope.blog.category = null;
    };
    $scope.blog = new BlogSrv({
        id: -1
    });

    $scope.save = function(){
        $scope.blog.$save(function(blog){
             $location.path('/viewblog/' + blog.id);
        });
    };    
    $scope.$on('$viewContentLoaded', function(){
        /*
        */
    });
}]);

admin.controller('ViewBlogCtrl', ['$scope', 'article', function($scope, article){
    $scope.article = article;
}]);

admin.controller('EditCtrl', ['$scope', '$location', 'article', function($scope, $location, article){
    $scope.article = article;

    $scope.save = function(){
        $scope.article.$save(function(article){
            $location.path('/view/' + article.id);
        });
    };

    $scope.remove = function(){
        $scope.article.$remove(function(article){
            $location.path('/');
        });
    };
}]);
