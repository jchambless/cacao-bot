import React from 'react';

interface StatusBadgeProps {
  status: 'online' | 'offline' | 'loading';
  label: string;
}

export const StatusBadge: React.FC<StatusBadgeProps> = ({ status, label }) => {
  const getStatusColor = () => {
    switch (status) {
      case 'online':
        return 'bg-green-100 text-green-800 border-green-200';
      case 'offline':
        return 'bg-red-100 text-red-800 border-red-200';
      case 'loading':
        return 'bg-yellow-100 text-yellow-800 border-yellow-200';
      default:
        return 'bg-gray-100 text-gray-800 border-gray-200';
    }
  };

  return (
    <span className={`inline-flex items-center px-3 py-1 rounded-full text-sm font-medium border ${getStatusColor()}`}>
      <span className={`w-2 h-2 rounded-full mr-2 ${status === 'online' ? 'bg-green-400' :
          status === 'offline' ? 'bg-red-400' :
            'bg-yellow-400'
        }`}></span>
      {label}
    </span>
  );
};
