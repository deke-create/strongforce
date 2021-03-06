using System;
using System.Collections.Generic;
using StrongForce.Core.Permissions;
using StrongForce.Core.Tests.Mocks;
using Xunit;

namespace StrongForce.Core.Tests
{
	public class ContractTests
	{
		private readonly BaseAddressFactory addressFactory = new RandomAddressFactory();

		[Fact]
		public void Receive_WhenPassedNull_ThrowsArgumentNullException()
		{
			var contract = StatefulObject.Create<FavoriteNumberContract>(new Dictionary<string, object>());
			var receiver = contract.RegisterWithRegistry(new InMemoryIntegration.FakeContractContext(this.addressFactory.CreateAddress()));
			Assert.Throws<ArgumentNullException>(() => receiver.Invoke(null));
		}

		[Fact]
		public void Receive_WhenReceivedSupportedAction_ReturnsTrue()
		{
			var contract = StatefulObject.Create<FavoriteNumberContract>(new Dictionary<string, object>() { { "User", null } });
			var receiver = contract.RegisterWithRegistry(new InMemoryIntegration.FakeContractContext(this.addressFactory.CreateAddress()));

			receiver.Invoke(new Message(
				contract.Address,
				contract.Address,
				contract.Address,
				SetFavoriteNumberAction.Type,
				new Dictionary<string, object>()
				{
					{ SetFavoriteNumberAction.Number, 0 },
				}));
		}
	}
}