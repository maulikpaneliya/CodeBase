// Express.js Basics - Web Application Framework

const express = require('express');
const cors = require('cors');
const morgan = require('morgan');
const helmet = require('helmet');
const rateLimit = require('express-rate-limit');

const app = express();
const port = process.env.PORT || 3000;

// 1. Middleware Configuration
// ------------------------

// Built-in middleware
app.use(express.json());  // Parse JSON bodies
app.use(express.urlencoded({ extended: true }));  // Parse URL-encoded bodies
app.use(express.static('public'));  // Serve static files

// Third-party middleware
app.use(cors());  // Enable CORS
app.use(morgan('dev'));  // HTTP request logger
app.use(helmet());  // Security headers

// Rate limiting
const limiter = rateLimit({
    windowMs: 15 * 60 * 1000,  // 15 minutes
    max: 100  // limit each IP to 100 requests per windowMs
});
app.use(limiter);

// Custom middleware
const requestLogger = (req, res, next) => {
    console.log(`${new Date().toISOString()} - ${req.method} ${req.url}`);
    next();
};
app.use(requestLogger);

// Authentication middleware
const authenticateUser = (req, res, next) => {
    const authHeader = req.headers.authorization;
    if (!authHeader) {
        return res.status(401).json({ error: 'Authorization header required' });
    }
    // Add authentication logic here
    next();
};

// 2. Route Handlers
// ---------------

// Basic routes
app.get('/', (req, res) => {
    res.send('Welcome to Express.js API');
});

// JSON response
app.get('/api/data', (req, res) => {
    res.json({
        message: 'Success',
        data: { key: 'value' }
    });
});

// Route parameters
app.get('/api/users/:id', (req, res) => {
    const userId = req.params.id;
    res.json({ userId });
});

// Query parameters
app.get('/api/search', (req, res) => {
    const { q, limit = 10, page = 1 } = req.query;
    res.json({ query: q, limit, page });
});

// POST request handling
app.post('/api/users', authenticateUser, (req, res) => {
    const userData = req.body;
    // Add user creation logic here
    res.status(201).json({
        message: 'User created',
        user: userData
    });
});

// 3. Router Module
// --------------

const userRouter = express.Router();

userRouter.get('/', (req, res) => {
    res.json({ message: 'Users list' });
});

userRouter.post('/', (req, res) => {
    res.json({ message: 'Create user' });
});

app.use('/api/users', userRouter);

// 4. Error Handling
// --------------

// 404 handler
app.use((req, res, next) => {
    res.status(404).json({ error: 'Not found' });
});

// Error handling middleware
app.use((err, req, res, next) => {
    console.error(err.stack);
    res.status(500).json({
        error: 'Internal server error',
        message: err.message
    });
});

// 5. File Upload Handling
// --------------------

const multer = require('multer');
const storage = multer.diskStorage({
    destination: (req, file, cb) => {
        cb(null, 'uploads/');
    },
    filename: (req, file, cb) => {
        cb(null, Date.now() + '-' + file.originalname);
    }
});
const upload = multer({ storage });

app.post('/api/upload', upload.single('file'), (req, res) => {
    res.json({
        message: 'File uploaded',
        file: req.file
    });
});

// 6. Template Engine Setup
// ---------------------

app.set('view engine', 'ejs');
app.set('views', './views');

app.get('/page', (req, res) => {
    res.render('index', {
        title: 'Express App',
        message: 'Welcome to Express'
    });
});

// 7. Database Connection
// -------------------

const mongoose = require('mongoose');

mongoose.connect(process.env.MONGODB_URI, {
    useNewUrlParser: true,
    useUnifiedTopology: true
})
.then(() => console.log('Connected to MongoDB'))
.catch(err => console.error('MongoDB connection error:', err));

// Start server
app.listen(port, () => {
    console.log(`Server running on port ${port}`);
});

// Best Practices:
// 1. Use appropriate middleware
// 2. Implement proper error handling
// 3. Structure routes logically
// 4. Use environment variables
// 5. Implement security measures
// 6. Handle async operations properly
// 7. Use proper HTTP status codes
// 8. Implement request validation
// 9. Use proper logging
// 10. Follow REST conventions
