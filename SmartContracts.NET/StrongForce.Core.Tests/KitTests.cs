using System;
using StrongForce.Core.Kits;
using StrongForce.Core.Permissions;
using StrongForce.Core.Tests.Mocks;
using Xunit;

namespace StrongForce.Core.Tests
{
	public class KitTests
	{
		[Fact]
		public void Kit_WhenInstantiated_CreatesContracts()
		{
			var favoriteContractsCount = 2;

			var factory = new SequentialAddressFactory();
			var registry = new ContractRegistry(factory);
			var kit = new FavoriteNumberKit(favoriteContractsCount);
			Address address = registry.CreateContract<KitContract>(kit);

			registry.HandleAction(address, new InstantiateKitAction(address));

			Assert.Equal(favoriteContractsCount + 1, factory.AddressCount);
		}

		[Fact]
		public void Kit_WhenInstantiatedTwice_Throws()
		{
			var registry = new ContractRegistry();
			var kit = new FavoriteNumberKit(2);
			Address address = registry.CreateContract<KitContract>(kit);

			registry.HandleAction(address, new InstantiateKitAction(address));
			Assert.Throws<InvalidOperationException>(() => registry.HandleAction(address, new InstantiateKitAction(address)));
		}
	}
}