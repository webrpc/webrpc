using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Interop.Tests
{
    // Check runs named test cases and collects their failures.
    internal sealed class Check
    {
        private readonly List<string> _failures = new List<string>();

        public IReadOnlyList<string> Failures => _failures;

        public async Task Run(string name, Func<Task> test)
        {
            try
            {
                await test();
                Console.WriteLine($"  ok    {name}");
            }
            catch (Exception ex)
            {
                _failures.Add($"{name}: {ex}");
                Console.WriteLine($"  FAIL  {name}: {ex.Message}");
            }
        }

        public void Run(string name, Action test)
        {
            Run(name, () =>
            {
                test();
                return Task.CompletedTask;
            }).GetAwaiter().GetResult();
        }
    }

    internal sealed class AssertException : Exception
    {
        public AssertException(string message) : base(message)
        {
        }
    }

    internal static class Assert
    {
        public static void Equal<T>(T expected, T actual, string what = "value")
        {
            if (!EqualityComparer<T>.Default.Equals(expected, actual))
            {
                throw new AssertException($"{what}: expected {Format(expected)}, got {Format(actual)}");
            }
        }

        public static void SequenceEqual<T>(IEnumerable<T> expected, IEnumerable<T>? actual, string what = "sequence")
        {
            if (actual == null || !expected.SequenceEqual(actual))
            {
                throw new AssertException($"{what}: expected [{string.Join(", ", expected.Select(x => Format(x)))}], got [{(actual == null ? "null" : string.Join(", ", actual.Select(x => Format(x))))}]");
            }
        }

        public static void DictionaryEqual<TKey, TValue>(IDictionary<TKey, TValue> expected, IDictionary<TKey, TValue>? actual, string what = "dictionary")
            where TKey : notnull
        {
            if (actual == null || actual.Count != expected.Count || expected.Any(kv => !actual.TryGetValue(kv.Key, out var v) || !EqualityComparer<TValue>.Default.Equals(kv.Value, v)))
            {
                throw new AssertException($"{what}: expected {{{string.Join(", ", expected.Select(kv => $"{kv.Key}: {Format(kv.Value)}"))}}}, got {(actual == null ? "null" : "{" + string.Join(", ", actual.Select(kv => $"{kv.Key}: {Format(kv.Value)}")) + "}")}");
            }
        }

        public static void True(bool condition, string what)
        {
            if (!condition)
            {
                throw new AssertException(what);
            }
        }

        public static async Task<TException> Throws<TException>(Func<Task> action, string what = "call")
            where TException : Exception
        {
            try
            {
                await action();
            }
            catch (TException ex)
            {
                return ex;
            }
            catch (Exception ex)
            {
                throw new AssertException($"{what}: expected {typeof(TException).Name}, got {ex.GetType().Name}: {ex.Message}");
            }
            throw new AssertException($"{what}: expected {typeof(TException).Name}, got no exception");
        }

        private static string Format<T>(T value)
        {
            return value switch
            {
                null => "null",
                string s => $"\"{s}\"",
                byte[] b => Convert.ToHexString(b),
                _ => value.ToString() ?? "",
            };
        }
    }
}
