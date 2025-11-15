const { execSync } = require('child_process');
const path = require('path');
const fs = require('fs');

// Get absolute path for .go-cache in the user-service directory
const serviceDir = path.resolve(__dirname, '..');
const cacheDir = path.resolve(serviceDir, '.go-cache');

// Ensure the directory exists
if (!fs.existsSync(cacheDir)) {
    fs.mkdirSync(cacheDir, { recursive: true });
}

// Set GOCACHE environment variable to absolute path
// Go requires an absolute path for GOCACHE
process.env.GOCACHE = cacheDir;

// Run go command
try {
    execSync('go run ./cmd/server', {
        stdio: 'inherit',
        cwd: serviceDir,
        env: {
            ...process.env,
            GOCACHE: cacheDir
        }
    });
} catch (error) {
    console.error('Error running go command:', error.message);
    process.exit(1);
}

