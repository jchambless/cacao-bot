import { useQuery } from '@tanstack/react-query';
import { healthCheck, getBotInfo, getServerStatus } from '../utils/api';

// Health check hook
export const useHealthCheck = () => {
  return useQuery({
    queryKey: ['health'],
    queryFn: healthCheck,
    refetchInterval: 30000, // Refetch every 30 seconds
  });
};

// Bot info hook
export const useBotInfo = () => {
  return useQuery({
    queryKey: ['botInfo'],
    queryFn: getBotInfo,
    staleTime: 10 * 60 * 1000, // 10 minutes
  });
};

// Server status hook
export const useServerStatus = () => {
  return useQuery({
    queryKey: ['serverStatus'],
    queryFn: getServerStatus,
    refetchInterval: 10000, // Refetch every 10 seconds
  });
};
