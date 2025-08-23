import { writable } from 'svelte/store';
import type { UserProfile, Car } from '../types';

// Default profile settings
const defaultProfile: UserProfile = {
	showSavedCars: true,
	savedCars: [],
	donationBannerDismissed: false
};

// Load profile from localStorage if available
function loadProfile(): UserProfile {
	if (typeof window === 'undefined') return defaultProfile;
	
	try {
		const stored = localStorage.getItem('car-search-profile');
		if (stored) {
			const parsed = JSON.parse(stored);
			return { ...defaultProfile, ...parsed };
		}
	} catch (error) {
		console.warn('Failed to load profile from localStorage:', error);
	}
	
	return defaultProfile;
}

// Save profile to localStorage
function saveProfile(profile: UserProfile) {
	if (typeof window === 'undefined') return;
	
	try {
		localStorage.setItem('car-search-profile', JSON.stringify(profile));
	} catch (error) {
		console.warn('Failed to save profile to localStorage:', error);
	}
}

// Create the profile store
function createProfileStore() {
	const { subscribe, set, update } = writable<UserProfile>(loadProfile());

	return {
		subscribe,
		set: (profile: UserProfile) => {
			set(profile);
			saveProfile(profile);
		},
		update: (updater: (profile: UserProfile) => UserProfile) => {
			update((profile) => {
				const newProfile = updater(profile);
				saveProfile(newProfile);
				return newProfile;
			});
		},
		toggleShowSavedCars: () => {
			update((profile) => {
				const newProfile = { ...profile, showSavedCars: !profile.showSavedCars };
				saveProfile(newProfile);
				return newProfile;
			});
		},
		saveCar: (car: Car) => {
			update((profile) => {
				// Avoid duplicates based on HSN and TSN
				const exists = profile.savedCars.some(
					(savedCar) => savedCar.HSN === car.HSN && savedCar.TSN === car.TSN
				);
				
				if (!exists) {
					const newProfile = {
						...profile,
						savedCars: [...profile.savedCars, car]
					};
					saveProfile(newProfile);
					return newProfile;
				}
				
				return profile;
			});
		},
		removeSavedCar: (car: Car) => {
			update((profile) => {
				const newProfile = {
					...profile,
					savedCars: profile.savedCars.filter(
						(savedCar) => !(savedCar.HSN === car.HSN && savedCar.TSN === car.TSN)
					)
				};
				saveProfile(newProfile);
				return newProfile;
			});
		},
		clearSavedCars: () => {
			update((profile) => {
				const newProfile = { ...profile, savedCars: [] };
				saveProfile(newProfile);
				return newProfile;
			});
		},
		dismissDonationBanner: () => {
			update((profile) => {
				const newProfile = { ...profile, donationBannerDismissed: true };
				saveProfile(newProfile);
				return newProfile;
			});
		}
	};
}

export const profileStore = createProfileStore();
