/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

app.controller('ListCtrl', ['$scope', 'links', function($scope, links){
    //console.log(links);
    $scope.links = links;
}]);

app.controller('ViewCtrl', ['$scope', '$location', 'link', function($scope, $location, link){
    $scope.link = link;

    $scope.edit = function(){
        $location.path('/edit/' + link.id);
    }
}]);

app.controller('EditCtrl', ['$scope', '$location', 'link', function($scope, $location, link){
    $scope.link = link;

    $scope.save = function(){
        $scope.link.$save(function(link){
            $location.path('/view/' + link.id);
        });
    };

    $scope.remove = function(){
        $scope.link.$remove(function(link){
            $location.path('/');
        });
    };
}]);

app.controller('NewCtrl', ['$scope', '$location', 'LinkSrv', function($scope, $location, LinkSrv){
    $scope.link = new LinkSrv({
        id: -1
    });

    $scope.save = function(){
        $scope.link.$save(function(link){
            $location.path('/view/' + link.id);
        });
    };
}]);

