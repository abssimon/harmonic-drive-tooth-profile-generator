"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var interp1_1 = __importDefault(require("../src/interp1"));
describe('interp1', function () {
    it('interp1 exists', function () {
        expect(interp1_1.default).toBeTruthy();
    });
    it('throws an error if the number of sample points does not equal the number of values', function () {
        expect(function () {
            interp1_1.default([1, 2, 3], [1, 2, 3, 4], [2, 2.5, 3]);
        }).toThrowError();
    });
    it('throws an error if two values of independent variable occur more than once', function () {
        expect(function () {
            interp1_1.default([5, 5, 8], [8, -3, 0], [6]);
        }).toThrowError();
    });
    it('throws an error if query points lie outside of sample points range', function () {
        expect(function () {
            interp1_1.default([-5, 0, 7], [3, 9, -8], [-6]);
        }).toThrowError();
        expect(function () {
            interp1_1.default([2, 4, 5], [-2, 4, -3], [7]);
        }).toThrowError();
    });
    it('returns an empty array for empty query points', function () {
        expect(interp1_1.default([], [], [])).toEqual([]);
        expect(interp1_1.default([7, 8, 9], [-5, 9, 81], [])).toEqual([]);
    });
    it('returns exact values', function () {
        expect(interp1_1.default([3], [7], [3])).toEqual([7]);
        expect(interp1_1.default([-8, 0, 5], [1, 92, 4], [-8, 5, 0])).toEqual([1, 4, 92]);
    });
    it('interpolates values linearly', function () {
        expect(interp1_1.default([-7, -3, 0, 4, 9], [0, 8, -4, -2, 3], [-7, -5, -3, -1.5, 0, 2, 4, 6.5, 9], 'linear')).toEqual([0, 4, 8, 2, -4, -3, -2, 0.5, 3]);
    });
    it('interpolates values with nearest neighbours', function () {
        expect(interp1_1.default([-5, 6, 8], [0, 4, -2], [-5, 0, 6, 6.99, 7, 8], 'nearest')).toEqual([0, 0, 4, 4, -2, -2]);
    });
    it('interpolates values with previous neighbours', function () {
        expect(interp1_1.default([-9, -7, 0, 1, 2], [0, 7, 3, -4, -2], [-9, -8, -7.1, -7, -1, 0, 1, 1.75, 2], 'previous')).toEqual([0, 0, 0, 7, 7, 3, -4, -4, -2]);
    });
    it('interpolates values with next neighbours', function () {
        expect(interp1_1.default([-4, -1, 0, 4, 7, 9], [-9, 0, 3, -4, 2, -1], [-4, -3.9, -1, -0.6, 0, 3, 4, 6, 7, 8, 9], 'next')).toEqual([-9, 0, 0, 3, 3, -4, -4, 2, 2, -1, -1]);
    });
    it('interpolates values linearly by default', function () {
        var xs = [-3, 0, 4, 5];
        var vs = [9, -2, 0, 3];
        var xqs = [-3, -1.5, 0, 2, 4, 5];
        expect(interp1_1.default(xs, vs, xqs)).toEqual(interp1_1.default(xs, vs, xqs));
    });
    it('does not require the sample points to be in order', function () {
        expect(interp1_1.default([9, -2, 0, 3], [8, 0, 3, 1], [-2, 0, 1.5, 3], 'linear')).toEqual([0, 3, 2, 1]);
    });
});
