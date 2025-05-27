import { useQuery } from '@tanstack/react-query';
import { healthCheck, getBotInfo, getServerStatus, getStats, getGuilds, getGuildInfo } from '../utils/api';

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

export const useStats = () => {
  return useQuery({
    queryKey: ['stats'],
    queryFn: getStats,
    refetchInterval: 60000, // Refetch every 60 seconds
  });
};

export const useGuilds = () => {
  return useQuery({
    queryKey: ['guilds'],
    queryFn: getGuilds,
    staleTime: 10 * 60 * 1000, // 10 minutes
  });
};

export const useGuildInfo = (guildId: string) => {
  return useQuery({
    queryKey: ['guildInfo', guildId],
    queryFn: () => getGuildInfo(guildId),
    enabled: !!guildId, // Only run if guildId is provided
    staleTime: 10 * 60 * 1000, // 10 minutes
  });
};