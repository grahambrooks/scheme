angular.module('viewApi', [])
    .controller('ViewController', ['$scope', '$http', '$location', function ($scope, $http, $location) {
        var params = {};
        window.location.search.replace(
            new RegExp("([^?=&]+)(=([^&]*))?", "g"),
            function ($0, $1, $2, $3) {
                params[$1] = $3;
            }
        );

        $scope.api = "";
        console.log("params", params);

        var query = '/api/apis/' + params.id;

        $http({
            method: 'GET',
            url: query
        }).then(function successCallback(response) {
            $scope.api = response.data;
        }, function errorCallback(response) {
            console.log("error from API " + response)
        });
    }]);