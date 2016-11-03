/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var app = angular.module('section', ['ngRoute', 'section.services', 'section.directives']);

app.config(['$interpolateProvider', function($interpolateProvider){
    $interpolateProvider.startSymbol('[[').endSymbol(']]');
}]);

app.config(['$routeProvider', function($routeProvider){
    $routeProvider.when('/', {
        controller: 'ListCtrl',
        resolve: {
            sections: function(MultiSectionLoader) {
                return MultiSectionLoader();
            }
        },
        templateUrl: '/app/section/views/list.html'
    }).when('/edit/:sectionId', {
        controller: 'EditCtrl',
    resolve: {
        section: function(SectionLoader){
            return SectionLoader();
        }
    },
    templateUrl: '/app/section/views/form.html'
    }).when('/view/:sectionId', {
        controller: 'ViewCtrl',
    resolve: {
        section: function(SectionLoader) {
            return SectionLoader();
        }
    },
    templateUrl: '/app/section/views/view.html'
    }).when('/new', {
        controller: 'NewCtrl',
    templateUrl: '/app/section/views/form.html'
    }).otherwise({redirectTo: '/'});
}]);
