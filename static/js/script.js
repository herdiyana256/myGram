function getData() {
    fetch('/photos', {
        method: 'GET',
        headers: {
            'Authorization': 'xxxxxxxxx' // token JWT
        }
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to fetch data');
        }
        return response.json();
    })
    .then(data => {
        console.log(data);
        //  manipulasi  dengan data yang diterima dari server
    })
    .catch(error => {
        console.error('Error fetching data:', error);
    });
}
