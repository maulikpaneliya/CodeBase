// Node.js File System Operations

const fs = require('fs').promises;
const path = require('path');

// 1. File Operations
// ----------------

// Read file
async function readFileExample() {
    try {
        const data = await fs.readFile('example.txt', 'utf8');
        console.log('File contents:', data);
    } catch (err) {
        console.error('Error reading file:', err);
    }
}

// Write file
async function writeFileExample() {
    try {
        await fs.writeFile('output.txt', 'Hello, World!', 'utf8');
        console.log('File written successfully');
    } catch (err) {
        console.error('Error writing file:', err);
    }
}

// Append to file
async function appendFileExample() {
    try {
        await fs.appendFile('log.txt', `\nLog entry: ${new Date().toISOString()}`);
        console.log('Data appended to file');
    } catch (err) {
        console.error('Error appending to file:', err);
    }
}

// 2. Directory Operations
// --------------------

// Create directory
async function createDirectoryExample() {
    try {
        await fs.mkdir('new-directory', { recursive: true });
        console.log('Directory created');
    } catch (err) {
        console.error('Error creating directory:', err);
    }
}

// Read directory contents
async function readDirectoryExample() {
    try {
        const files = await fs.readdir('.');
        console.log('Directory contents:', files);
    } catch (err) {
        console.error('Error reading directory:', err);
    }
}

// 3. File Stats
// -----------

async function getFileStatsExample() {
    try {
        const stats = await fs.stat('example.txt');
        console.log('File size:', stats.size);
        console.log('Is file:', stats.isFile());
        console.log('Is directory:', stats.isDirectory());
        console.log('Last modified:', stats.mtime);
    } catch (err) {
        console.error('Error getting file stats:', err);
    }
}

// 4. Working with Paths
// ------------------

function pathOperationsExample() {
    const filePath = '/path/to/file/example.txt';
    
    console.log('Directory name:', path.dirname(filePath));
    console.log('Base name:', path.basename(filePath));
    console.log('Extension:', path.extname(filePath));
    
    const newPath = path.join(__dirname, 'files', 'example.txt');
    console.log('Joined path:', newPath);
}

// 5. File Streams
// ------------

const { createReadStream, createWriteStream } = require('fs');

// Read stream
function readStreamExample() {
    const readStream = createReadStream('large-file.txt', {
        encoding: 'utf8',
        highWaterMark: 64 * 1024 // 64KB chunks
    });

    readStream.on('data', (chunk) => {
        console.log('Received chunk:', chunk.length, 'bytes');
    });

    readStream.on('end', () => {
        console.log('Finished reading file');
    });

    readStream.on('error', (err) => {
        console.error('Error reading file:', err);
    });
}

// Write stream
function writeStreamExample() {
    const writeStream = createWriteStream('output-stream.txt');

    writeStream.write('Hello\n');
    writeStream.write('World\n');
    writeStream.end();

    writeStream.on('finish', () => {
        console.log('Finished writing to file');
    });

    writeStream.on('error', (err) => {
        console.error('Error writing to file:', err);
    });
}

// Copy file using streams
function copyFileStreamExample() {
    const readStream = createReadStream('source.txt');
    const writeStream = createWriteStream('destination.txt');

    readStream.pipe(writeStream);

    writeStream.on('finish', () => {
        console.log('File copied successfully');
    });
}

// 6. Recursive Directory Operations
// -----------------------------

async function recursiveDirectoryExample() {
    async function processDirectory(dirPath) {
        try {
            const entries = await fs.readdir(dirPath, { withFileTypes: true });

            for (const entry of entries) {
                const fullPath = path.join(dirPath, entry.name);

                if (entry.isDirectory()) {
                    await processDirectory(fullPath);
                } else {
                    const stats = await fs.stat(fullPath);
                    console.log(`File: ${fullPath}, Size: ${stats.size} bytes`);
                }
            }
        } catch (err) {
            console.error('Error processing directory:', err);
        }
    }

    await processDirectory('.');
}

// 7. Watch File Changes
// ------------------

function watchFileExample() {
    const watcher = fs.watch('watched-file.txt', (eventType, filename) => {
        console.log(`File ${filename} changed: ${eventType}`);
    });

    // Stop watching after 60 seconds
    setTimeout(() => {
        watcher.close();
        console.log('Stopped watching file');
    }, 60000);
}

// 8. Temporary Files
// ---------------

const os = require('os');

async function tempFileExample() {
    try {
        const tempDir = os.tmpdir();
        const tempFilePath = path.join(tempDir, 'temp-' + Date.now());
        
        await fs.writeFile(tempFilePath, 'Temporary data');
        console.log('Temp file created:', tempFilePath);
        
        // Clean up after 5 seconds
        setTimeout(async () => {
            await fs.unlink(tempFilePath);
            console.log('Temp file deleted');
        }, 5000);
    } catch (err) {
        console.error('Error handling temp file:', err);
    }
}

// Best Practices:
// 1. Always use async/await with fs.promises
// 2. Handle errors properly
// 3. Use streams for large files
// 4. Clean up temporary files
// 5. Close file handles when done
// 6. Use appropriate file permissions
// 7. Validate file paths
// 8. Use path module for cross-platform compatibility
// 9. Implement proper error logging
// 10. Consider performance with large directories
