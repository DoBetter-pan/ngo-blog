/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

app.controller('ListCtrl', ['$scope', 'sections', function($scope, sections){
    //console.log(sections);
    $scope.sections = sections;
}]);

app.controller('ViewCtrl', ['$scope', '$location', 'section', function($scope, $location, section){
    $scope.section = section;

    $scope.edit = function(){
        $location.path('/edit/' + section.id);
    }
}]);

app.controller('EditCtrl', ['$scope', '$location', 'section', function($scope, $location, section){
    $scope.section = section;

    $scope.save = function(){
        $scope.section.$save(function(section){
            $location.path('/view/' + section.id);
        });
    };

    $scope.remove = function(){
        $scope.section.$remove(function(section){
            $location.path('/');
        });
    };
}]);

app.controller('NewCtrl', ['$scope', '$location', 'SectionSrv', function($scope, $location, SectionSrv){
    $scope.section = new SectionSrv({
        id: -1
    });

    $scope.save = function(){
        $scope.section.$save(function(section){
            $location.path('/view/' + section.id);
        });
    };
}]);

