document.addEventListener('DOMContentLoaded', function () {
    const fetchItemA = document.getElementById('fetchItemA');
    const fetchItemB = document.getElementById('fetchItemB');
    const responseData = document.getElementById('responseData');

    function fetchData(apiEndpoint) {
        fetch(`http://localhost:3000${apiEndpoint}`)
            .then(response => {
                if (!response.ok) {
                    throw new Error(`Network response was not ok (status: ${response.status})`);
                }
                return response.json();
            })
            .then(data => {
                responseData.textContent = JSON.stringify(data, null, 2);
                responseData.style.display = 'block';
            })
            .catch(error => {
                console.error('There was a problem with the fetch operation:', error);
            });
    }

    fetchItemA.addEventListener('click', function () {
        fetchData('/getItemA');
    });

    fetchItemB.addEventListener('click', function () {
        fetchData('/getItemB');
    });
});
