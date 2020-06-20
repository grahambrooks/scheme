angular.module('SchemeApp', [])
    .controller('SearchController', ['$scope', '$http', function ($scope, $http) {
        $scope.searchText = "";
        $scope.searchResponse = "waiting";
        $scope.resultCount = 0;
        $scope.errorResponse = null;

        $scope.updateResults = function () {
            var query = '/api/search?query=' + $scope.searchText;

            $http({
                method: 'GET',
                url: query
            }).then(function successCallback(response) {
                $scope.errorResponse = null;
                $scope.searchResponse = response.data;
                $scope.resultCount = $scope.searchResponse.hits.total.value
            }, function errorCallback(response) {
                console.log("error from API " + response);
                $scope.errorResponse = response.data;
            });
        };

        $scope.updateResults()
    }]);