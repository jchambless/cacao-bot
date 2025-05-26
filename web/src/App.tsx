import React from 'react';
import { useHealthCheck } from './hooks/useApi';
import { StatusBadge } from './components/StatusBadge';

function App() {
  const { data: healthData, isLoading, error } = useHealthCheck();

  const getHealthStatus = () => {
    if (isLoading) return 'loading';
    if (error || !healthData?.ok) return 'offline';
    return 'online';
  };

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header className="bg-white shadow-sm border-b">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center py-6">
            <div className="flex items-center">
              <h1 className="text-3xl font-bold text-gray-900">Cacao Bot</h1>
              <span className="ml-4 text-sm text-gray-500">Discord Bot Dashboard</span>
            </div>
            <StatusBadge
              status={getHealthStatus()}
              label={getHealthStatus() === 'online' ? 'Bot Online' : getHealthStatus() === 'loading' ? 'Checking...' : 'Bot Offline'}
            />
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">

          {/* Bot Status Card */}
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-lg font-semibold text-gray-900 mb-4">Bot Status</h3>
            <div className="space-y-3">
              <div className="flex justify-between items-center">
                <span className="text-sm text-gray-600">Health Check</span>
                <StatusBadge status={getHealthStatus()} label={getHealthStatus()} />
              </div>
              {healthData && (
                <div className="text-xs text-gray-500">
                  Last checked: {new Date().toLocaleTimeString()}
                </div>
              )}
            </div>
          </div>

          {/* Server Commands Card */}
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-lg font-semibold text-gray-900 mb-4">Quick Commands</h3>
            <div className="space-y-2">
              <button className="w-full px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
                Server Status
              </button>
              <button className="w-full px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 transition-colors">
                Player List
              </button>
              <button className="w-full px-4 py-2 bg-yellow-600 text-white rounded-md hover:bg-yellow-700 transition-colors">
                Bot Info
              </button>
            </div>
          </div>

          {/* Statistics Card */}
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-lg font-semibold text-gray-900 mb-4">Statistics</h3>
            <div className="space-y-3">
              <div className="flex justify-between">
                <span className="text-sm text-gray-600">Guilds</span>
                <span className="text-sm font-medium">--</span>
              </div>
              <div className="flex justify-between">
                <span className="text-sm text-gray-600">Uptime</span>
                <span className="text-sm font-medium">--</span>
              </div>
              <div className="flex justify-between">
                <span className="text-sm text-gray-600">Commands Run</span>
                <span className="text-sm font-medium">--</span>
              </div>
            </div>
          </div>

        </div>

        {/* Available Commands Section */}
        <div className="mt-8">
          <div className="bg-white rounded-lg shadow">
            <div className="px-6 py-4 border-b border-gray-200">
              <h3 className="text-lg font-semibold text-gray-900">Available Commands</h3>
            </div>
            <div className="p-6">
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                {[
                  { name: 'help', description: 'Get help about bot commands' },
                  { name: 'ban', description: 'Ban a player from the server' },
                  { name: 'kick', description: 'Kick a player from the server' },
                  { name: 'list', description: 'Show online players' },
                  { name: 'ping', description: 'Check server status' },
                  { name: 'about', description: 'Get bot information' },
                ].map((command) => (
                  <div key={command.name} className="border border-gray-200 rounded-lg p-4">
                    <h4 className="font-medium text-gray-900 mb-1">!mc {command.name}</h4>
                    <p className="text-sm text-gray-600">{command.description}</p>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
}

export default App;
