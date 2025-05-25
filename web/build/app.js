// Main React app entry point

// Wait for DOM to be fully loaded
document.addEventListener('DOMContentLoaded', () => {
    // Create root container for React
    const rootElement = document.getElementById('root');
    
    if (!rootElement) {
        console.error('Root element not found. Make sure you have a div with id="root" in your HTML.');
        return;
    }

    // App component
    const App = () => {
        const [count, setCount] = React.useState(0);
        
        return React.createElement('div', { 
            className: 'min-h-screen bg-gray-100 flex flex-col items-center justify-center p-4'
        }, [
            React.createElement('div', { 
                key: 'container',
                className: 'bg-white shadow-md rounded-lg p-8 max-w-md w-full' 
            }, [
                React.createElement('h1', { 
                    key: 'title',
                    className: 'text-3xl font-bold text-center text-blue-600 mb-6' 
                }, 'Cacao Bot'),
                
                React.createElement('p', { 
                    key: 'description',
                    className: 'text-gray-700 text-center mb-6' 
                }, 'Welcome to Cacao Bot! This is a simple React application with Tailwind CSS.'),
                
                React.createElement('div', { 
                    key: 'counter',
                    className: 'flex flex-col items-center gap-4'
                }, [
                    React.createElement('p', { 
                        key: 'count',
                        className: 'text-xl text-gray-800' 
                    }, `Count: ${count}`),
                    
                    React.createElement('button', { 
                        key: 'button',
                        className: 'bg-blue-500 hover:bg-blue-600 text-white font-medium py-2 px-4 rounded transition-colors',
                        onClick: () => setCount(count + 1)
                    }, 'Increment')
                ])
            ])
        ]);
    };

    // Render the app to the DOM
    ReactDOM.render(React.createElement(App), rootElement);
});