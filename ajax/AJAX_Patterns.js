// AJAX Examples and Patterns
// This file contains various AJAX implementation examples using different methods

// 1. XMLHttpRequest (Vanilla JavaScript)
// Basic GET Request
function makeGETRequest(url) {
    const xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);
    
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                console.log(JSON.parse(xhr.responseText));
            } else {
                console.error('Error:', xhr.status);
            }
        }
    };
    
    xhr.send();
}

// POST Request with XMLHttpRequest
function makePOSTRequest(url, data) {
    const xhr = new XMLHttpRequest();
    xhr.open('POST', url, true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                console.log(JSON.parse(xhr.responseText));
            } else {
                console.error('Error:', xhr.status);
            }
        }
    };
    
    xhr.send(JSON.stringify(data));
}

// 2. Fetch API (Modern JavaScript)
// Basic GET Request
async function fetchGET(url) {
    try {
        const response = await fetch(url);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error:', error);
    }
}

// POST Request with Fetch
async function fetchPOST(url, data) {
    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data)
        });
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const result = await response.json();
        return result;
    } catch (error) {
        console.error('Error:', error);
    }
}

// 3. jQuery AJAX
// Basic GET Request
$.ajax({
    url: 'api/endpoint',
    method: 'GET',
    dataType: 'json',
    success: function(response) {
        console.log(response);
    },
    error: function(xhr, status, error) {
        console.error('Error:', error);
    }
});

// POST Request with jQuery
$.ajax({
    url: 'api/endpoint',
    method: 'POST',
    data: JSON.stringify({ key: 'value' }),
    contentType: 'application/json',
    success: function(response) {
        console.log(response);
    },
    error: function(xhr, status, error) {
        console.error('Error:', error);
    }
});

// 4. Axios (Modern HTTP Client)
// Basic GET Request
axios.get('api/endpoint')
    .then(response => {
        console.log(response.data);
    })
    .catch(error => {
        console.error('Error:', error);
    });

// POST Request with Axios
axios.post('api/endpoint', {
    key: 'value'
})
    .then(response => {
        console.log(response.data);
    })
    .catch(error => {
        console.error('Error:', error);
    });

// 5. File Upload Example
function uploadFile(file) {
    const formData = new FormData();
    formData.append('file', file);
    
    fetch('api/upload', {
        method: 'POST',
        body: formData
    })
        .then(response => response.json())
        .then(data => console.log(data))
        .catch(error => console.error('Error:', error));
}

// 6. Progress Tracking
function uploadWithProgress(file) {
    const xhr = new XMLHttpRequest();
    const formData = new FormData();
    formData.append('file', file);
    
    xhr.upload.onprogress = function(e) {
        if (e.lengthComputable) {
            const percentComplete = (e.loaded / e.total) * 100;
            console.log(percentComplete + '% uploaded');
        }
    };
    
    xhr.onload = function() {
        if (xhr.status === 200) {
            console.log('Upload complete');
        }
    };
    
    xhr.open('POST', 'api/upload', true);
    xhr.send(formData);
}

// 7. Handling CORS
fetch('https://api.example.com/data', {
    method: 'GET',
    credentials: 'include', // Include cookies
    headers: {
        'Authorization': 'Bearer token'
    }
})
    .then(response => response.json())
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));

// 8. Abort Controller
function fetchWithTimeout(url, timeout = 5000) {
    const controller = new AbortController();
    const signal = controller.signal;
    
    setTimeout(() => controller.abort(), timeout);
    
    return fetch(url, { signal })
        .then(response => response.json())
        .catch(error => {
            if (error.name === 'AbortError') {
                throw new Error('Request timed out');
            }
            throw error;
        });
}

// 9. Error Handling Patterns
async function robustFetch(url, options = {}) {
    const MAX_RETRIES = 3;
    let attempt = 0;
    
    while (attempt < MAX_RETRIES) {
        try {
            const response = await fetch(url, options);
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return await response.json();
        } catch (error) {
            attempt++;
            if (attempt === MAX_RETRIES) throw error;
            await new Promise(resolve => setTimeout(resolve, 1000 * attempt));
        }
    }
}
