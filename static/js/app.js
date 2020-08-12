'use strict';

angular.module('tour', ['ui', 'tour.services', 'tour.controllers', 'tour.directives', 'tour.values', 'ng']).

config(['$routeProvider', '$locationProvider',
    function($routeProvider, $locationProvider) {
        $routeProvider.
        when('/', {
            redirectTo: '/list'
        }).
        when('/list', {
            templateUrl: '/static/partials/list.html',
        }).
        when('/:lessonId/:pageNumber', {
            templateUrl: '/static/partials/editor.html',
            controller: 'EditorCtrl'
        }).
        when('/:lessonId', {
            redirectTo: '/:lessonId/1'
        }).
        otherwise({
            redirectTo: '/'
        });

        $locationProvider.html5Mode(true).hashPrefix('!');
    }
]);

// handle mapping from old paths (#42) to the new organization.
// run(function($rootScope, $location, mapping) {
//     $rootScope.$on( "$locationChangeStart", function(event, next) {
//         var url = document.createElement('a');
//         url.href = next;
//         if (url.pathname != '/' || url.hash == '') {
//             return;
//         }
//         $location.hash('');
//         var m = mapping[url.hash];
//         if (m === undefined) {
//             console.log('unknown url, redirecting home');
//             $location.path('/list');
//             return;
//         }
//         $location.path(m);
//     });
// });
