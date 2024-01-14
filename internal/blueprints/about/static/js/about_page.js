const date = {
    qqq: "1111",
};

fetch('/about', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify(date), // Convert the object to a JSON string
})
    .then((response) => response.json())
    .then((data) => {
        console.log('data:', data);
    })
    .catch((error) => {
        console.log('Error:', error);
    });